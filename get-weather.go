package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
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
	var endPoint string = "https://www.spesolution.com/trial/get-weather"
	var bearer = "Bearer " + <ACCESS TOKEN HERE>
	var jsonStr = []byte(`
	{
    		"resource_type" : "get_weather"
	}`)

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
    request.Header.Set("x-api-key", "dHJpYXNiM2V1bTpueXhxZFBpOTJYWWFDTTEw")

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