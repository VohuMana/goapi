package structbuilder

import (
	"testing"
	"encoding/json"
	"strings"
)

func compareMembers(expected, returned []string, t *testing.T) {
	if len(expected) != len(returned) {
		t.Fatalf("Returned member size was not equal, expected %v. Got %v", len(expected), len(returned))
	}

	for index := range expected {
		if strings.Compare(returned[index], expected[index]) == 0 {
			t.Fatalf("Member strings do not match expected %v. Got %v", expected[index], returned[index])
		}
	}
}

func TestSimpleJson(t *testing.T) {
	simpleJson := "{ \"testKey\": \"testString\" }"
	expectedMembers := []string {"TestKey string `json:testKey`"}
	jsonObject := make(map[string]interface{})
	err := json.Unmarshal([]byte(simpleJson), &jsonObject)
	if err != nil {
		t.Fatal(err)
	}

	structs := GenerateStructs(jsonObject, false)

	for _,v := range structs {
		compareMembers(expectedMembers, v, t)
	}
}

func TestMultipleJsonFields(t *testing.T) {
	simpleJson := "{ \"testKey\": \"testString\", \"odin\":123, \"frey\":true, \"thor\":[1,2,3,4] }"
	expectedMembers := []string {"TestKey string `json:testKey`", "Odin float64 `json:odin`", "Frey bool `json:frey`", "Thor []float64 `json:thor`"}
	jsonObject := make(map[string]interface{})
	err := json.Unmarshal([]byte(simpleJson), &jsonObject)
	if err != nil {
		t.Fatal(err)
	}

	structs := GenerateStructs(jsonObject, false)

	for _,v := range structs {
		compareMembers(expectedMembers, v, t)
	}
}

func TestMultipleObjects(t *testing.T) {
	simpleJson := "{ \"testKey\": \"testString\", \"odin\":123, \"frey\":true, \"thor\":[1,2,3,4], \"loki\":{\"fenrir\":\"wolf\",\"hel\":[\"helheim\"]} }"

	jsonObject := make(map[string]interface{})
	err := json.Unmarshal([]byte(simpleJson), &jsonObject)
	if err != nil {
		t.Fatal(err)
	}

	structs := GenerateStructs(jsonObject, false)

	if len(structs) != 2 {
		t.Fatalf("Expected number of objects to be 2, got %v", len(structs))
	}
}