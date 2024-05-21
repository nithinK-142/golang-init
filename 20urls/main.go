package main

import (
	"fmt"
	"net/url"
)

func main() {
	HandleGetRequests()
}

func HandleGetRequests() {
	const weburl = "https://nithin-me-git-preproduction-nithink142s-projects.vercel.app/"

	result, _ := url.Parse(weburl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)

	createUrl := &url.URL{
		Scheme: "https",
		Host:   "nithin-me-git-preproduction-nithink142s-projects.vercel.app",
		Path:   "/",
	}

	fmt.Println(createUrl.String())
}
