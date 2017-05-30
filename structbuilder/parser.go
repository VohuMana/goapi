package structbuilder

import (
	"strings"
	"fmt"
	"encoding/hex"
	"math/rand"
)

func generateStructName() string {
	randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return hex.EncodeToString(randBytes)
}

func parseArray(object []interface{}, structName string, structs map[string][]string) string {
	typeName, ok := getArrayType(object)
	if !ok {
		switch object[0].(type) {
			case []interface{}:
				typeName = "[]" + parseArray(object[0].([]interface{}), structName, structs)
			
			case map[string]interface{}:
				newStructName := generateStructName()
				typeName = "[]" + newStructName
				parseStructs(object[0].(map[string]interface{}), newStructName, structs)

			default:
				// TODO: Return an error here
				fmt.Println("Error unknown array type")
		}
	}

	return typeName
}

func parseStructs(object map[string]interface{}, name string, structs map[string][]string) {	
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
				valueType = parseArray(value.([]interface{}), name, structs)			
			
			case map[string]interface{}:
				valueType = generateStructName()
				parseStructs(value.(map[string]interface{}), valueType, structs)

			case nil:
				fmt.Printf("Key %v was nil, type is unknown", key)

			default:
				fmt.Printf("Don't know how to parse %v\n", key)
				continue
		}

		structs[name] = append(structs[name], fmt.Sprintf("%v %v `json:%v`\n", strings.Title(key), valueType, key))
	}
}

func getArrayType(arr []interface{}) (string, bool) {
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

// GenerateStructs generates the golang structs from a map.  The return value is a map of struct names to an array of struct members.
func GenerateStructs(object map[string]interface{}) map[string][]string{
	structs := make(map[string][]string)
	parseStructs(object, generateStructName(), structs)
	return structs
}