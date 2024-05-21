package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dice := rand.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("dice gave 1")
	case 2:
		fmt.Println("dice gave 2")
		fallthrough
	case 3:
		fmt.Println("dice gave 3")
	case 4:
		fmt.Println("dice gave 4")
	case 5:
		fmt.Println("dice gave 5")
	case 6:
		fmt.Println("dice gave 6")
	default:
		fmt.Println("what was that?")
	}
}
