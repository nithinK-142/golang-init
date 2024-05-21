package main

import "fmt"

func main() {
	// creating a pointer
	var ptr *int
	fmt.Println("default value of ptr", ptr)

	// reference to memory address
	value := 10
	num := &value
	fmt.Println("memory address reference of value", num)
	fmt.Println("value of value", *num)

	// manipulating pointer directly
	*num = *num * 10
	fmt.Println("after pointer manipulation", *num)
}
