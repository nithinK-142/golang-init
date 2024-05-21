package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./Go.txt")
	checkNilErr(err)

	// length, err := io.WriteString(file, "here we GO!!!")
	length, err := file.WriteString("here we GO!!!")
	checkNilErr(err)
	fmt.Println(length)

	defer file.Close()

	ReadFile("./Go.txt")
}

func ReadFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
