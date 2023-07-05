package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

func (c *APIClient) GetData() ([]byte, error) {
	url := c.baseURL + "/data"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request using the HTTP client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Return the response body
	return body, nil
}
func Function1() {
	// Create an instance of the API client
	client := APIClient{
		baseURL:    "http://api.example.com",
		httpClient: &http.Client{},
	}

	// Call the GetData method to retrieve data
	data, err := client.GetData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process the retrieved data
	fmt.Println("Data:", string(data))
}

func Function2() {
	// Create an instance of the API client
	client := APIClient{
		baseURL:    "http://api.example.com",
		httpClient: &http.Client{},
	}

	// Call the GetData method to retrieve data
	data, err := client.GetData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process the retrieved data
	fmt.Println("Data:", string(data))
}
