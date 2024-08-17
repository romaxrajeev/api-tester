package api

import (
	"api-tester/functions"
	"api-tester/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PerformApiTest(c *gin.Context) {
	headers := c.Request.Header
	requiredDataFromHeaders, err := utils.GetTestapifyData(&headers);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
		return
    }
	if requiredDataFromHeaders["Method"] == "GET" || requiredDataFromHeaders["Method"] == "OPTIONS"{
		response := functions.ExecuteRequest(requiredDataFromHeaders["Method"], requiredDataFromHeaders["Url"], nil)
        c.JSON(int(response.StatusCode), response.Data)
	} else {
		body := c.Request.Body
		response := functions.ExecuteRequest(requiredDataFromHeaders["Method"], requiredDataFromHeaders["Url"], body)
		c.JSON(int(response.StatusCode), response.Data)
	}
}