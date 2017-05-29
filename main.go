package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"math/rand"
	"encoding/hex"
)

func GenerateStructName() string {
	randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return hex.EncodeToString(randBytes)
}

func ParseArray(object []interface{}, structName string, structs map[string][]string) string {
	typeName, ok := GetArrayType(object)
	if !ok {
		switch object[0].(type) {
			case []interface{}:
				typeName = "[]" + ParseArray(object[0].([]interface{}), structName, structs)
			
			case map[string]interface{}:
				newStructName := GenerateStructName()
				typeName = "[]" + newStructName
				ParseStructs(object[0].(map[string]interface{}), newStructName, structs)

			default:
				// TODO: Return an error here
				fmt.Println("Error unknown array type")
		}
	}

	return typeName
}

func ParseStructs(object map[string]interface{}, name string, structs map[string][]string) {	
	structs[name] = []string{}

	for key,value := range object {
		valueType := ""

		switch value.(type) {
			case string:
				valueType = "string"

			case float64:
				valueType = "float64"

			case bool:
				valueType = "bool"
			
			case []interface{}:
				valueType = ParseArray(value.([]interface{}), name, structs)			
			
			case map[string]interface{}:
				valueType = GenerateStructName()
				ParseStructs(value.(map[string]interface{}), valueType, structs)

			default:
				fmt.Printf("Don't know how to parse %v\n", key)
				continue
		}

		structs[name] = append(structs[name], fmt.Sprintf("%v %v `json:%v`\n", key, valueType, key))
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
	structs := make(map[string][]string)
	ParseStructs(jsonObject, GenerateStructName(), structs)

	for structName, members := range structs {
		fmt.Printf("type %v struct {\n", structName)
		for _, member := range members {
			fmt.Printf("\t%v", member)
		}
		fmt.Println("}")
	}
}