package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/vohumana/goapi/structbuilder"
)


func main() {
	fileData, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON
	jsonObject := make(map[string]interface{})
	err = json.Unmarshal(fileData, &jsonObject)
	if err != nil {
		log.Fatal(err)
	}

	// Enumerate the interface for all the objects
	structs := structbuilder.GenerateStructs(jsonObject)

	for structName, members := range structs {
		fmt.Printf("type %v struct {\n", structName)
		for _, member := range members {
			fmt.Printf("\t%v", member)
		}
		fmt.Println("}")
	}
}