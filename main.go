package main

import (
	"api-tester/api"

	"github.com/gin-gonic/gin"
)

func main(){
	// method := "get"
	// url := "https://jsonplaceholder.typicode.com/users/1"

	// response := functions.ExecuteRequest(method, "",  url)
	// fmt.Println(response.Data) 
	router := gin.Default()
    router.POST("/api/perform-api-test", api.PerformApiTest)

    router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    // router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for linux/macos)
    // router.Run(":8081") // listen and serve on 0.0.0.0:8081 (for other systems)
    // router.RunTLS(":8081", "./cert.pem", "./key.pem")
}