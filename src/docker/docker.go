package docker

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/orisano/uds"
)

const (
	dockerSocketPath = "/var/run/docker.sock"
)

// Client provides methods for accessing docker host
type Client struct {
	unixHTTPClient *http.Client
}

// FileInfo specifies the file name and its content
type FileInfo struct {
	Name string
	Text string
}

// Init starts a http socket client for the unix domain socket interface
func (c *Client) Init() {
	c.unixHTTPClient = uds.NewClient(dockerSocketPath)
}

// IsConnected checks if the connection was established
func (c *Client) IsConnected() bool {
	return c.unixHTTPClient != nil
}

// CreateImage creates a docker image with the received files, returns err if exist
func (c *Client) CreateImage(name string, files ...FileInfo) (err error) {
	tarBuffer := bytes.Buffer{}
	tarWriter := tar.NewWriter(&tarBuffer)

	for _, file := range files {
		tarHeader := &tar.Header{Name: file.Name, Mode: 0600, Size: int64(len(file.Text))}
		tarWriter.WriteHeader(tarHeader)
		tarWriter.Write([]byte(file.Text))
	}
	tarWriter.Close()

	response, err := c.unixHTTPClient.Post(
		fmt.Sprintf("http://docker/build?t=%v", name),
		"application/x-tar",
		&tarBuffer,
	)

	io.Copy(os.Stdout, response.Body)
	response.Body.Close()

	return err
}

// CreateContainer initializes a container with the received image, returns the container id and error if exist
func (c *Client) CreateContainer(image string) (string, error) {
	fmt.Println("CreateContainer")
	response, err := c.unixHTTPClient.Post(
		"http://docker/containers/create",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf(`{ "Image": "%v" }`, image))),
	)
	fmt.Println("Err")
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	body, err2 := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println("body")
	fmt.Println(body)
	fmt.Println("err2")
	fmt.Println(err2)
	containerID := string(body[7:71])
	fmt.Println("containerID")
	fmt.Println(containerID)
	return containerID, err2
}

// StartContainer starts the container with the received containerID, returns the containerIP and error if exist
func (c *Client) StartContainer(containerID string) (string, error) {
	response, err := c.unixHTTPClient.Post(
		fmt.Sprintf("http://docker/containers/%v/start", containerID),
		"application/json",
		nil,
	)
	
	
	fmt.Println("response")
	fmt.Println(response)
	
	fmt.Println("err")
	fmt.Println(err)

	io.Copy(os.Stdout, response.Body)
	response.Body.Close()

	inspectResponse, _ := c.unixHTTPClient.Get(
		fmt.Sprintf("http://docker/containers/%v/json", containerID),
	)
	
	fmt.Println("inspectResponse")
	fmt.Println(inspectResponse)
	
	
	var inspectJSON map[string]interface{}
	json.NewDecoder(inspectResponse.Body).Decode(&inspectJSON)
	containerIP := inspectJSON["NetworkSettings"].(map[string]interface{})["Networks"].(map[string]interface{})["bridge"].(map[string]interface{})["IPAddress"].(string)

	
	
	fmt.Println("containerIP")
	fmt.Println(containerIP)
	
	
	return containerIP, err
}

// StopContainer stops the container with the received container Id, returns error if exist
func (c *Client) StopContainer(containerID string) error {
	response, err := c.unixHTTPClient.Post(
		fmt.Sprintf("http://docker/containers/%v/kill", containerID),
		"application/json",
		nil,
	)
	io.Copy(os.Stdout, response.Body)
	response.Body.Close()

	return err
}

// DeleteContainer deletes the container with the received container Id, returns error if exist
func (c *Client) DeleteContainer(containerID string) error {
	request, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("http://docker/containers/%v", containerID),
		nil,
	)
	response, err := c.unixHTTPClient.Do(request)
	io.Copy(os.Stdout, response.Body)
	response.Body.Close()

	return err
}

// DeleteImage deletes the image with the received name, returns error if exist
func (c *Client) DeleteImage(name string) error {
	request, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("http://docker/images/%v", name),
		nil,
	)
	response, err := c.unixHTTPClient.Do(request)
	io.Copy(os.Stdout, response.Body)
	response.Body.Close()

	return err
}
