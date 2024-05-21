package main

import "fmt"

func main() {
	// Immediately invoked anonymous function
	func() {
		fmt.Println("autoGreetings")
	}()

	geetings()

	result := add(3, 10)
	fmt.Println(result)

	fmt.Println(proAdder(1, 2, 3, 4, 5))
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// anonymous function
var add = func(x, y int) int {
	return x + y
}

func geetings() {
	fmt.Println("Namaste Go lang!")
}
