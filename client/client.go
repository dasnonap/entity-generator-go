package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestMethod int

const (
	GET RequestMethod = iota
	POST
	PUT
	DELETE
	OPTIONS
)

var requestMethodName = map[RequestMethod]string{
	GET:     "GET",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	OPTIONS: "OPTIONS",
}

func (method RequestMethod) String() string {
	return requestMethodName[method]
}

type Client struct {
	Method RequestMethod
	Url    string
	Body   []byte
}

type Response struct {
	Headers    http.Header `json:"headers"`
	StatusCode int `json:"statusCode"`
	Body       bytes.Buffer `json:"body"`
}

func NewClient(method RequestMethod, url string) Client {
	client := Client{}

	client.Method = method
	client.Url = url
	client.Body = nil

	return client
}

// Execute client request
func (client Client) SendRequest() (Response, error) {
	request, err := http.NewRequest(client.Method.String(), client.Url, nil)
	
	if (err != nil) {
		return Response{}, fmt.Errorf("error while creating a request %v", err)
	}
	
	handlerClient := http.Client{}
	response, err := handlerClient.Do(request)

	if (err != nil) {
		return Response{}, fmt.Errorf("error while sending request %v", err)
	}

	defer response.Body.Close()
	
	var data bytes.Buffer

	// Collect Body data
	// err = json.NewDecoder(response.Body).Decode(&data)
	// rawBody, err = 
	if (err != nil) {
		return Response{}, fmt.Errorf("error while decoding request %s", err)
	}

	return Response{
		response.Header,
		response.StatusCode,
		data,
	}, nil
}

// Prettier response
func (response Response) PrettyPrint() {
	jsonData, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}