package functions

import (
	"api-tester/structs"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpGet(url string, data *bytes.Reader) structs.Response {
    client := &http.Client{}
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return structs.Response{
            StatusCode: uint16(http.StatusInternalServerError),
            Data: fmt.Sprintf("There is an error while making request : %v", err),
        }
    }

    resp, err := client.Do(req)
    if err != nil {
        return structs.Response{
            StatusCode: uint16(http.StatusInternalServerError),
            Data: fmt.Sprintf("There is an error while executing request : %v", err),
        }
    }

    if resp.StatusCode != http.StatusOK {
        return structs.Response{
            StatusCode: uint16(resp.StatusCode),
            Data: fmt.Sprintf("There is an error in statuscode : %v", err),
        }
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return structs.Response{
            StatusCode: uint16(http.StatusInternalServerError),
            Data: fmt.Sprintf("There is an error in reading body : %v", err),
        }
    }

    return structs.Response{
        StatusCode: 200,
        Data: string(bodyBytes),
    }
}