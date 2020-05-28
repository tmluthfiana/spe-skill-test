package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func main() {
	var endPoint string = "https://www.spesolution.com/trial/secret-key"

	var jsonStr = []byte(`
	{
    		"Username" : "triasb3eum",
		"Password": "nyxqdPi92XYaCM10"
	}`)

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// use httpClient to send request
	response, err := httpClient.Do(req)
	if err != nil && response == nil {
		fmt.Printf("Error sending request to API endpoint. %+v", err)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()

		// Let's check if the work actually is done
		// We have seen inconsistencies even when we get 200 OK response
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Couldn't parse response body. %+v", err)
		}

		fmt.Println("Response Body:", string(body))
	}

}
