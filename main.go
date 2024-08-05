package main

import (
	"api-tester/functions"
	"fmt"
)

func main(){
	method := "get"
	url := "https://jsonplaceholder.typicode.com/users/1"

	response := functions.ExecuteRequest(method, "", url)
	fmt.Println(response.Data) 
}