package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

func ParseStructs(object map[string]interface{}) {
	for key,value := range object {
		valueType := ""

		// TODO: Add arrays and subobjects
		switch value.(type) {
			case string:
				valueType = "string"

			case float64:
				valueType = "float64"

			case bool:
				valueType = "bool"
			
			case []interface{}:
				ok := true
				valueType, ok = GetArrayType(value.([]interface{}))
				if !ok {
					switch value.([]interface{})[0].(type) {
						case map[string]interface{}:
							valueType = "Foo"
							ParseStructs(value.([]interface{})[0].(map[string]interface{}))
					}
				}
			
			case map[string]interface{}:
				ParseStructs(value.(map[string]interface{}))
			default:
				fmt.Printf("Don't know how to parse %v\n", key)
				continue
		}

		fmt.Printf("%v %v `json:%v`\n", key, valueType, key)
	}
}

func GetArrayType(arr []interface{}) (string, bool) {
	switch arr[0].(type) {
		case string:
			return "[]string", true

		case float64:
			return "[]float64", true

		case bool:
			return "[]bool", true

		// Failed to parse because it is either not one of the above types, or is an array of arrays, or is an array of objects
		default:
			return "", false
	}
}

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
	fmt.Println(jsonObject)

	ParseStructs(jsonObject)
}