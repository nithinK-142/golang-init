package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "user input"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter somethng...")

	// comma ok || err
	inp, _ := reader.ReadString('\n')
	fmt.Println(inp)
}
