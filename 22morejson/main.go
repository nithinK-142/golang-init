package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Tags     []string `json:"tags"`
}

func main() {
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	fmt.Println("Encode json")
	products := []Product{
		{"Mechanical Keyboard", 999, "Amazon", []string{"smartphone", "apple", "premium"}},
		{"PlayStation 5", 499, "Flipkart", []string{"gaming", "sony", "nextgen"}},
		{"MacBook", 2499, "Ebay", []string{"laptop", "apple", "pro", "developer"}},
	}

	finalJson, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	fmt.Println("\nDecode json")
	jsonDataFromWeb := []byte(`{
	"name": "Mechanical Keyboard",
	"price": 999,
	"platform": "Amazon",
	"tags": ["smartphone","apple","premium"]}`)

	var product Product

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("valid json")
		json.Unmarshal(jsonDataFromWeb, &product)
		fmt.Printf("%#v\n", product)
	} else {
		fmt.Println("invalid json")
	}

	var jsonData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &jsonData)
	fmt.Printf("%#v\n", jsonData)

	fmt.Println()
	for i, j := range jsonData {
		fmt.Printf("%v : %v - %T\n", i, j, j)
	}
}
