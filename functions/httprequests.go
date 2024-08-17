package functions

import (
	"api-tester/structs"
	"api-tester/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpGet (url string, data *bytes.Reader) structs.Response {
    client := &http.Client{}
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return utils.ReturnErrorObject(fmt.Sprintf("There is an error while making request : %v", err.Error()), http.StatusInternalServerError)
    }

    resp, err := client.Do(req)
    fmt.Println(resp)
    if err != nil {
        return utils.ReturnErrorObject(fmt.Sprintf("There is an error while executing request : %v", err.Error()), http.StatusInternalServerError)
    }

    if resp.StatusCode != http.StatusOK {
        return utils.ReturnErrorObject(fmt.Sprintf("There is an error in statuscode : %v", resp.Body), uint16(resp.StatusCode))
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return utils.ReturnErrorObject(fmt.Sprintf("There is an error in reading body : %v", err.Error()), http.StatusInternalServerError)
    }

    body, err := json.Marshal(bodyBytes)
    if err != nil {
        return utils.ReturnErrorObject(fmt.Sprintf("There is an error in reading body : %v", err.Error()), http.StatusInternalServerError)
    }

    return structs.Response{
        StatusCode: 200,
        Data: body,
    }
}