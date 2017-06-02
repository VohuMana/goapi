package filebuilder

import (
	"fmt"
	"bytes"
)

func BuildFile(packageName string, structs map[string][]string) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("package %v\n\n", packageName))

	for structName, members := range structs {
		buffer.WriteString(fmt.Sprintf("type %v struct {\n", structName))
		for _, member := range members {
			buffer.WriteString(fmt.Sprintf("\t%v", member))
		}
		buffer.WriteString(fmt.Sprintf("}\n\n"))
	}

	return buffer.String()
}