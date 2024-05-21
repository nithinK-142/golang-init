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
	const url = "https://nithin-me-git-preproduction-nithink142s-projects.vercel.app/"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("status code", response.StatusCode)

	databytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(databytes))

	// we have to mannually close connection
	defer response.Body.Close()
}
