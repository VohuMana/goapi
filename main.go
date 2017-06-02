package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"flag"

	"github.com/vohumana/goapi/structbuilder"
	"github.com/vohumana/goapi/filebuilder"
)


func main() {
	var (
		packageName = flag.String("packagename", "api", "Name of the package used in generated file.")
		fileName = flag.String("filename", "generated_api.go", "File name/path to output generated structs to.")
		inputFileName = flag.String("inputfile", "", "Name of the JSON file to parse the API from.")
		useKeyNames = flag.Bool("usekeynames", false, "Add this to use key names for struct names.  If this is not specified it will generate a random hex string.")
	)

	flag.Parse()

	fileData, err := ioutil.ReadFile(*inputFileName)
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
	structs := structbuilder.GenerateStructs(jsonObject, *useKeyNames)
	apiContract := filebuilder.BuildFile(*packageName, structs)

	// Write to disk
	err = ioutil.WriteFile(*fileName, []byte(apiContract), 0)
	if err != nil {
		log.Fatal(err)
	}
}