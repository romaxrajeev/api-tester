package functions

import (
	"api-tester/structs"
	"bytes"
	"fmt"
	"net/http"
)

// Get the parameters passed via headers
func GetParametersFromData(){

}

// Use the parameters to send an API request to that server
func ExecuteRequest(method string, data string, url string) structs.Response{
	// Get the string to be a byte array
	body := []byte(data)
	bodyReader := bytes.NewReader(body)
	// Switch case
	switch method {
	case "get":
		response := HttpGet(url, bodyReader)
		return response
	default:
		fmt.Println("Invalid Http method")
		return structs.Response{
			StatusCode: http.StatusInternalServerError,
			Data: "Invalid HTTP method",
		}
	}
}