package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ricardorobson/sl-handler/src/database"
	"github.com/ricardorobson/sl-handler/src/docker"
)

var (
	db                            = database.Database{}
	dockerClient                  = docker.Client{}
	dockerfile, _                 = ioutil.ReadFile("schemas/js/Dockerfile")
	serverJS, _                   = ioutil.ReadFile("schemas/js/server.js")
)

const (
	functionEndpoint = "/function/"
	callEndpoint     = "/call/"
	port             = ":80"
	timeOutSeconds   = 60
)

func main() {
	db.Connect()
	dockerClient.Init()
	if isConnected := dockerClient.IsConnected(); !isConnected {
		fmt.Println("Failed to connect")
	}

	http.HandleFunc(functionEndpoint, function)
	http.HandleFunc(callEndpoint, call)
	http.ListenAndServe(port, nil)
}

func function(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		functionGet(res, req)
	case "POST":
		functionPost(res, req)
	case "DELETE":
		functionDelete(res, req)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func functionGet(res http.ResponseWriter, req *http.Request) {
	var argument = req.RequestURI[len(functionEndpoint):]
	if !strings.EqualFold(argument, "") {
		var function = functionGetByName(argument)
		if function == "" {
			res.Write([]byte(fmt.Sprintf("Function with name %v not found", argument)))
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.Write([]byte(function))

	} else {
		var functions = functionGetAll()
		res.Write([]byte(functions))
	}
}

func functionGetAll() string {
	return string(db.SelectAllFunction())
}

func functionGetByName(argument string) string {
	return string(db.SelectFunction(argument))
}

func functionPost(res http.ResponseWriter, req *http.Request) {
	name, code, pack := ExtractFunction(res, req.Body)

	if len(db.SelectFunction(name)) == 0 {
		dockerClient.CreateImage(
			name,
			docker.FileInfo{Name: "Dockerfile", Text: string(dockerfile)},
			docker.FileInfo{Name: "server.js", Text: string(serverJS)},
			docker.FileInfo{Name: "package.json", Text: pack},
			docker.FileInfo{Name: "code.js", Text: code},
		)
		db.InsertFunction(name, code, pack)
		var function = functionGetByName(name)
		res.Write([]byte(function))
		res.Write([]byte(fmt.Sprintf("Function Created at %v%v\n", req.RequestURI, name)))
		res.WriteHeader(http.StatusCreated)
	} else {
		http.Error(res, "Function already exist\n"+http.StatusText(http.StatusConflict), http.StatusConflict)
	}
}

func ExtractFunction(res http.ResponseWriter, jsonBodyReq io.Reader) (name string, code, pack string) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["name"].(string), bodyData["code"].(string), bodyData["package"].(string)
}

func functionDelete(res http.ResponseWriter, req *http.Request) {
	var name = strings.Split(req.RequestURI, "/")[2]

	if len(db.SelectFunction(name)) > 0 {
		dockerClient.DeleteImage(name)
		var success = db.DeleteFunction(name)
		if !success {
			res.Write([]byte(fmt.Sprintf("Cannot Delete function %v\n", name)))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Write([]byte(fmt.Sprintf("Function Deleted [%v] %v\n", req.Method, req.RequestURI)))
		res.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(res, "Function don't exist\n"+http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func call(res http.ResponseWriter, req *http.Request) {
	requestData := req.RequestURI[6:]
	slashIndex := strings.Index(requestData, "/")
	if slashIndex == -1 {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Function endpoint not provided"))
		return
	}
	imageName := requestData[:slashIndex]
	fmt.Printf("imageName")
	fmt.Printf(imageName)

	containerID, err := dockerClient.CreateContainer(imageName)
	fmt.Printf("## Container ID: %v\n", containerID)

	containerIP, containerStartTime := dockerClient.StartContainer(containerID)
	fmt.Printf("## Container IP: %v\n", containerIP)
	fmt.Printf("## Start Container Time: %v\n", containerStartTime)

	startApplicationConnectionTime := time.Now()
	var applicationRunTime time.Duration
	gatewayReq, err := http.NewRequest(req.Method, fmt.Sprintf("http://%v:8080/%v", containerIP, requestData[len(imageName)+1:]), req.Body)
	var gatewayRes *http.Response
	for true {
		startApplicationRunTime := time.Now()
		gatewayRes, err = http.DefaultClient.Do(gatewayReq)
		if err == nil {
			applicationRunTime = time.Since(startApplicationRunTime)
			fmt.Printf("## Request Run Time: %v\n", applicationRunTime)
			break
		}
		applicationConnectionTime := time.Since(startApplicationConnectionTime)

		if float64(applicationConnectionTime)*0.000000001 >= timeOutSeconds {
			fmt.Printf("## Request Timeout Fail to %v\n", containerIP)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	applicationConnectionTime := time.Since(startApplicationConnectionTime)
	fmt.Printf("## Request Time: %v\n", applicationConnectionTime)

	applicationCode := gatewayRes.StatusCode
	applicationBody, _ := ioutil.ReadAll(gatewayRes.Body)
	res.WriteHeader(applicationCode)
	res.Write(applicationBody)

	dockerClient.StopContainer(containerID)
	dockerClient.DeleteContainer(containerID)
}
