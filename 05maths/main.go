package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Random number from crypto/rand
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number:", randomNum)

	// Rounding and truncating
	x := 3.14159
	fmt.Println("Round 3.14159 to nearest integer:", math.Round(x))
	fmt.Println("Floor of 3.14159:", math.Floor(x))
	fmt.Println("Ceiling of 3.14159:", math.Ceil(x))
}
