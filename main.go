package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("simple.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	
	fmt.Println("JsonFile Opened")

	// Parse JSON
	
	// Enumerate the interface for all the objects

}