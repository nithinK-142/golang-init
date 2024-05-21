package main

import "fmt"

func main() {
	nit := User{"Nithin", "nithi@dev.com", true, 25}
	fmt.Println(nit)
	fmt.Printf("%+v", nit)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
