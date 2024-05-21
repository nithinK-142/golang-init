package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	HandleGetRequests()
	HandlePostRequests()
	HandlePostFormRequests()
}

func HandleGetRequests() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(databytes))
}

func HandlePostRequests() {
	const url = "http://localhost:8000/post"

	reqbody := strings.NewReader(`{
		"foo":"bar",
		"fizz":"buzz"
	}`)

	response, err := http.Post(url, "application/json", reqbody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(databytes))
}

func HandlePostFormRequests() {
	const serverurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("name", "nit")
	data.Add("git", "nithinK-142")

	response, err := http.PostForm(serverurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
