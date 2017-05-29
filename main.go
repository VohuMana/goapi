package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	fileData, err := ioutil.ReadFile("simple.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileData))

	// Parse JSON
	jsonObject := make(map[string]interface{})
	err = json.Unmarshal(fileData, &jsonObject)
	if err != nil {
		log.Fatal(err)
	}

	// Enumerate the interface for all the objects
	structs := []string{}
	fmt.Println(jsonObject)
	for key,value := range jsonObject {
		valueType := ""

		// TODO: Add arrays and subobjects
		switch value.(type) {
			case string:
				valueType = "string"

			case float64:
				valueType = "float64"

			case bool:
				valueType = "bool"

			default:
				fmt.Printf("Don't know how to parse %v\n", key)
				continue
		}

		structs = append(structs, fmt.Sprintf("%v %v `json:%v`", key, valueType, key))		
	}

	for _,v := range structs {
		fmt.Println(v)
	}
}