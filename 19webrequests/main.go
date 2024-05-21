package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	HandleGetRequests()
}

func HandleGetRequests() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code", response.StatusCode)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
