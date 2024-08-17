package functions

import (
	"api-tester/structs"
	"api-tester/utils"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Get the parameters passed via headers
func GetParametersFromData(){

}

// Use the parameters to send an API request to that server
func ExecuteRequest(method string, url string, data io.ReadCloser) structs.Response{
	// Get the string to be a byte array
	var body []byte; 
	var err error = nil;

	if(data == nil){
		body = []byte{};
	} else{
		body, err = io.ReadAll(data)
		if err != nil {
			fmt.Println("Error reading data")
			return utils.ReturnErrorObject("Error reading data", http.StatusInternalServerError)
		}
	}
	// Switch case
	switch method {
	case "GET":
		response := HttpGet(url, bytes.NewReader(body))
		return response
	default:
		fmt.Println("Invalid Http method")
		return utils.ReturnErrorObject("Invalid Http method", http.StatusBadRequest)
	}
}