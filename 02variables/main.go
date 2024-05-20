package main

import "fmt"

const Value = "here we GO"

func main() {
	// default values
	var str string
	fmt.Println("default values")
	fmt.Println(str)

	var smallInt int8
	fmt.Println(smallInt)

	var smallFloat float32
	fmt.Println(smallFloat)

	// variable types
	fmt.Println("\nvariable types")
	fmt.Printf("%T\n", str)

	fmt.Printf("%T\n", smallInt)

	fmt.Printf("%T\n", smallFloat)

	// assigning values
	fmt.Println("\nassigning values")
	str = "GO"
	fmt.Println(str)

	smallInt = 127
	fmt.Println(smallInt)

	smallFloat = 255.23080823
	fmt.Println(smallFloat)

	// implicit types
	fmt.Println("\nimplicit type example")
	var str2 = "GO2"
	fmt.Printf("%T\n", str2)
	fmt.Println(str2)

	// walrus operator
	// allowed inside a method only
	fmt.Println("\nwalrus operator")
	num := 100
	fmt.Printf("%T\n", num)
	fmt.Println(num)

	// public variable
	fmt.Println("\npublic variable")
	fmt.Println(Value)
}
