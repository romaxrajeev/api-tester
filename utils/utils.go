package utils

import (
	"api-tester/structs"
	"fmt"
	"net/http"
)

func GetTestapifyData(header *http.Header) (map[string]string, error) {
	testApifyHeaders := make(map[string]string)
	for key, value := range *header {
		switch key {
			case "Testapify-Token":
                testApifyHeaders["Token"] = value[0]
            case "Testapify-Url":
                testApifyHeaders["Url"] = value[0]
            case "Testapify-Method":
                testApifyHeaders["Method"] = value[0]
            default:
				// Do nothing
		}
    }
	if len(testApifyHeaders) != 3 {
        return nil, fmt.Errorf("all required headers are not present")
    }
	return testApifyHeaders, nil
}

func ReturnErrorObject(errorMessage string, statusCode uint16) structs.Response{
	var sendStatusCode uint16
	if statusCode == 0 {
		sendStatusCode = http.StatusInternalServerError
	} else {
		sendStatusCode = statusCode
	}
	return structs.Response{
		StatusCode: sendStatusCode,
		Data: []byte(fmt.Sprintf("Error from Testapify: %s", errorMessage)),
	}
}