package main

import "fmt"

func main() {
	// defer working(LIFO)
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Go")
	printNum()
}

func printNum() {
	for i := 1; i < 6; i++ {
		defer fmt.Println(i)
	}
}
