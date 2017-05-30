package main

import (
	"encoding/json"
	"log"
	"io/ioutil"

	"github.com/vohumana/goapi/structbuilder"
	"github.com/vohumana/goapi/filebuilder"
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

	// Generate the API contract
	structs := structbuilder.GenerateStructs(jsonObject)
	apiContract := filebuilder.BuildFile("api", structs)

	// Write to disk
	err = ioutil.WriteFile("output.go", []byte(apiContract), 0)
	if err != nil {
		log.Fatal(err)
	}
}