package builder

import (
	"fmt"
	"strings"
)

const structTemplateString = `
type %s struct {
	%s
}
`

const typeTemplateString = `
type %s %s
`

type structBuilder struct {
	structs map[string]Type
}

var processed []string

func inArray(array []string, target string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

func (builder structBuilder) Build() string {
	goFileContent := "package types\n"
	for _, item := range builder.structs {
		//if inArray(processed, item.Name) {
		//	continue
		//}

		switch item.Type {
		case "struct":
			goFileContent += fmt.Sprintf(structTemplateString, strings.ReplaceAll(item.Name, "-", ""), strings.Join(item.Fields, "\n\t"))
		case "type":
			goFileContent += fmt.Sprintf(typeTemplateString, strings.ReplaceAll(item.Name, "-", ""), item.TypeName)
		}
		println(item.Name)
		processed = append(processed, item.Name)
	}

	return goFileContent
}

func (builder structBuilder) AddType(structName, itemType, itemTypeName string) {
	builder.structs[structName] = Type{
		structName,
		itemType,
		itemTypeName,
		[]string{},
	}
}

func (builder structBuilder) AddField(structName string, fieldName, fieldType string, isArray bool) {
	if isArray {
		fieldType = "[]" + fieldType
	}
	structItem, _ := builder.structs[structName]
	structItem.Fields = append(
		structItem.Fields,
		fmt.Sprintf("%s %s `xml:\"%s\" json:\"%s\"`", strings.ReplaceAll(strings.Title(fieldName), "-", ""), strings.ReplaceAll(fieldType, "-", ""), fieldName, fieldName),
	)
	builder.structs[structName] = structItem
}

func (builder structBuilder) AddCharDataField(structName string, fieldType string) {
	structItem, _ := builder.structs[structName]
	structItem.Fields = append(
		structItem.Fields,
		fmt.Sprintf("%s %s `xml:\",chardata\" json:\"%s\"`", "Value", strings.ReplaceAll(fieldType, "-", ""), "value"),
	)
	builder.structs[structName] = structItem
}

func (builder structBuilder) AddParent(structName string, parentStructName string) {
	structItem, _ := builder.structs[structName]
	structItem.Fields = append(
		structItem.Fields,
		parentStructName,
	)
	builder.structs[structName] = structItem
}

func (builder structBuilder) AddAttribute(structName string, fieldName, fieldType string) {
	structItem, _ := builder.structs[structName]
	structItem.Fields = append(
		structItem.Fields,
		fmt.Sprintf("%s %s `xml:\"%s,attr\" json:\"%s\"`", strings.ReplaceAll(strings.Title(fieldName), "-", ""), strings.ReplaceAll(fieldType, "-", ""), fieldName, fieldName),
	)
	builder.structs[structName] = structItem
}

func CreateStructType() TypeBuilder {
	return structBuilder{structs: make(map[string]Type)}
}
