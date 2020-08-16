package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
)

const (
	trigonometryEndpoint = "/trigonometry/"
	port                 = ":8080"
)

var trigonometryStruct = new(Trigonometry)

func main() {
	http.HandleFunc(trigonometryEndpoint, trigonometry)

	http.ListenAndServe(port, nil)
}

func trigonometry(res http.ResponseWriter, req *http.Request) {
	requestData := strings.Split(req.RequestURI, "/")[2]
	switch requestData {
	case "Sin":
		x := extractTrigonometrySin(res, req.Body)
		result := trigonometryStruct.Sin(x)
		res.Write([]byte(fmt.Sprintf(fmt.Sprintf("%f", result))))
		res.WriteHeader(http.StatusCreated)
	case "Cos":
		x := extractTrigonometryCos(res, req.Body)
		result := trigonometryStruct.Cos(x)
		res.Write([]byte(fmt.Sprintf(fmt.Sprintf("%f", result))))
		res.WriteHeader(http.StatusCreated)
	case "Tan":
		x := extractTrigonometryTan(res, req.Body)
		result := trigonometryStruct.Tan(x)
		res.Write([]byte(fmt.Sprintf(fmt.Sprintf("%f", result))))
		res.WriteHeader(http.StatusCreated)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func extractTrigonometrySin(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["x"].(float64)
}

func extractTrigonometryCos(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["x"].(float64)
}

func extractTrigonometryTan(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["x"].(float64)
}

type Trigonometry struct{}

func (t *Trigonometry) Sin(x float64) (result float64) {
	return math.Sin(x)
}

func (t *Trigonometry) Cos(x float64) (result float64) {
	return math.Cos(x)
}

func (t *Trigonometry) Tan(x float64) (result float64) {
	return math.Tan(x)
}

//http://www.unit-conversion.info/texttools/replace-text/
//https://www.gillmeister-software.com/online-tools/text/remove-line-breaks.aspx
/*

package main\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"io\"\n\t\"math\"\n\t\"net/http\"\n\t\"strings\"\n)\nconst (\n\ttrigonometryEndpoint = \"/trigonometry/\"\n\tport                 = \":8080\"\n)\nvar trigonometryStruct = new(Trigonometry)\nfunc main() {\n\thttp.HandleFunc(trigonometryEndpoint, trigonometry)\n\thttp.ListenAndServe(port, nil)\n}\nfunc trigonometry(res http.ResponseWriter, req *http.Request) {\n\trequestData := strings.Split(req.RequestURI, \"/\")[2]\n\tswitch requestData {\n\tcase \"Sin\":\n\t\tx := extractTrigonometrySin(res, req.Body)\n\t\tresult := trigonometryStruct.Sin(x)\n\t\tres.Write([]byte(fmt.Sprintf(fmt.Sprintf(\"%f\", result))))\n\t\tres.WriteHeader(http.StatusCreated)\n\tcase \"Cos\":\n\t\tx := extractTrigonometryCos(res, req.Body)\n\t\tresult := trigonometryStruct.Cos(x)\n\t\tres.Write([]byte(fmt.Sprintf(fmt.Sprintf(\"%f\", result))))\n\t\tres.WriteHeader(http.StatusCreated)\n\tcase \"Tan\":\n\t\tx := extractTrigonometryTan(res, req.Body)\n\t\tresult := trigonometryStruct.Tan(x)\n\t\tres.Write([]byte(fmt.Sprintf(fmt.Sprintf(\"%f\", result))))\n\t\tres.WriteHeader(http.StatusCreated)\n\tdefault:\n\t\thttp.Error(res, \"Method not allowed\", http.StatusMethodNotAllowed)\n\t}\n}\nfunc extractTrigonometrySin(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {\n\tvar jsonBody interface{}\n\terr := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)\n\tif err != nil {\n\t\thttp.Error(res, err.Error(), 400)\n\t\treturn\n\t}\n\tvar bodyData = jsonBody.(map[string]interface{})\n\treturn bodyData[\"x\"].(float64)\n}\nfunc extractTrigonometryCos(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {\n\tvar jsonBody interface{}\n\terr := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)\n\tif err != nil {\n\t\thttp.Error(res, err.Error(), 400)\n\t\treturn\n\t}\n\tvar bodyData = jsonBody.(map[string]interface{})\n\treturn bodyData[\"x\"].(float64)\n}\nfunc extractTrigonometryTan(res http.ResponseWriter, jsonBodyReq io.Reader) (x float64) {\n\tvar jsonBody interface{}\n\terr := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)\n\tif err != nil {\n\t\thttp.Error(res, err.Error(), 400)\n\t\treturn\n\t}\n\tvar bodyData = jsonBody.(map[string]interface{})\n\treturn bodyData[\"x\"].(float64)\n}\ntype Trigonometry struct{}\nfunc (t *Trigonometry) Sin(x float64) (result float64) {\n\treturn math.Sin(x)\n}\nfunc (t *Trigonometry) Cos(x float64) (result float64) {\n\treturn math.Cos(x)\n}\nfunc (t *Trigonometry) Tan(x float64) (result float64) {\n\treturn math.Tan(x)\n}
*/
