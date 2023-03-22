package xsdProcessor

import (
	"encoding/xml"
	"fmt"
	"main/elements"
	"main/xsdProcessor/builder"
	"os"
	"path/filepath"
	"strings"
)

type XsdProcessor interface {
	Process(location string)
}

type xsdProcessor struct {
	Prefix           string
	StoragePath      string
	SchemaFetcher    SchemaFetcher
	typeBuilder      builder.TypeBuilder
	processedSchemas []string
	processedFiles   []string
}

func Create(prefix, storagePath string) XsdProcessor {
	return xsdProcessor{
		Prefix:        prefix,
		StoragePath:   storagePath,
		SchemaFetcher: XMLFetcher{},
	}
}

func (processor xsdProcessor) Process(location string) {
	processor.typeBuilder = builder.CreateStructType()
	file, err := processor.SchemaFetcher.Fetch(location)
	dir, fileName := filepath.Split(location)
	fmt.Printf("Parsing: %s\n", location)

	if err != nil {
		fmt.Printf("Failed to read XSD file: %v\n", err)
		os.Exit(1)
	}
	var schema elements.Schema
	err = xml.Unmarshal(file, &schema)
	if err != nil {
		println(err.Error())
	}
	processor.processSchema(schema)
	err = os.WriteFile(processor.StoragePath+"/"+fileName[:len(fileName)-len(filepath.Ext(fileName))]+".go", []byte(processor.typeBuilder.Build()), 0644)
	if err != nil {
		println(err.Error())
	}
	for _, include := range schema.Includes {
		schemaLocation := include.SchemaLocation
		schemaDir, _ := filepath.Split(schemaLocation)
		if schemaDir == "" {
			schemaLocation = dir + schemaLocation
		}

		processor.Process(schemaLocation)
	}
}

func (processor xsdProcessor) addProcessedSchema(targetSchema string) bool {
	for _, processedSchema := range processor.processedSchemas {
		if targetSchema == processedSchema {
			return false
		}
	}

	processor.processedSchemas = append(processor.processedSchemas, targetSchema)
	return true
}

func (processor xsdProcessor) processSimpleTypeRestriction(structName string, elementName string, restriction elements.SimpleTypeRestriction) {
	processor.typeBuilder.AddField(structName, elementName, processor.getType(restriction.Base), false)
}

func (processor xsdProcessor) processSchema(schema elements.Schema) {
	processor.processElements(nil, schema.Elements)
	for _, simpleType := range schema.SimpleTypes {
		processor.processSimpleType(nil, nil, simpleType)
	}
	for _, complexType := range schema.ComplexTypes {
		processor.processComplexType(nil, complexType)
	}
}

func (processor xsdProcessor) processList(structName, elementName string, list elements.List) {
	processor.typeBuilder.AddField(structName, elementName, processor.getType(list.ItemType), false)
}

func (processor xsdProcessor) processUnion(structName, elementName string, union elements.Union) {
	if union.MemberTypes != nil {
		processor.typeBuilder.AddField(structName, elementName, processor.getType(strings.Split(*union.MemberTypes, " ")[0]), false)
	} else {
		if len(union.SimpleTypes) > 0 {
			processor.processSimpleType(&structName, &elementName, union.SimpleTypes[0])
		} else {
			println("Can't find types of union, skipped")
		}
	}
}

func (processor xsdProcessor) processSimpleType(structName *string, elementName *string, simpleType elements.SimpleType) {
	var name string
	if elementName != nil {
		name = *elementName
	} else if simpleType.Name != nil {
		name = *simpleType.Name
	} else {
		return
	}
	if structName == nil {
		if simpleType.Restriction != nil {
			processor.typeBuilder.AddType(name, "type", processor.getType(simpleType.Restriction.Base))
		} else if simpleType.List != nil {
			processor.typeBuilder.AddType(name, "type", processor.getType(simpleType.List.ItemType))
		} else {
			if simpleType.Union.MemberTypes != nil {
				processor.typeBuilder.AddType(name, "type", processor.getType(strings.Split(*simpleType.Union.MemberTypes, " ")[0]))
			} else {
				if len(simpleType.Union.SimpleTypes) > 0 {
					processor.processSimpleType(structName, elementName, simpleType.Union.SimpleTypes[0])
				} else {
					println("Can't find types of union, skipped")
				}
			}
		}
	} else {
		if simpleType.Restriction != nil {
			processor.processSimpleTypeRestriction(*structName, name, *simpleType.Restriction)
		} else if simpleType.List != nil {
			processor.processList(*structName, name, *simpleType.List)
		} else {
			processor.processUnion(*structName, name, *simpleType.Union)
		}
	}
}

func (processor xsdProcessor) processAll(structName string, all elements.All) {
	processor.processElements(&structName, all.Elements)
}

func (processor xsdProcessor) processChoice(structName string, choice elements.Choice) {
	processor.processElements(&structName, choice.Elements)
	for _, group := range choice.Groups {
		processor.processGroup(structName, group)
	}
	for _, choiceItem := range choice.Choices {
		processor.processChoice(structName, choiceItem)
	}
	for _, sequence := range choice.Sequences {
		processor.processSequence(structName, sequence)
	}
}

func (processor xsdProcessor) processSequence(structName string, sequence elements.Sequence) {
	processor.processElements(&structName, sequence.Elements)
	for _, group := range sequence.Groups {
		processor.processGroup(structName, group)
	}
	for _, choice := range sequence.Choices {
		processor.processChoice(structName, choice)
	}
	for _, sequenceItem := range sequence.Sequences {
		processor.processSequence(structName, sequenceItem)
	}
}

func (processor xsdProcessor) processGroup(structName string, group elements.Group) {
	if group.All != nil {
		processor.processAll(structName, *group.All)
	} else if group.Choice != nil {
		processor.processChoice(structName, *group.Choice)
	} else {
		processor.processSequence(structName, *group.Sequence)
	}
}

func (processor xsdProcessor) processAttributes(structName string, attributes []elements.Attribute) {
	for _, attribute := range attributes {
		if attribute.Name != nil {
			if attribute.SimpleType != nil {
				if attribute.SimpleType.Restriction != nil {
					processor.typeBuilder.AddAttribute(structName, *attribute.Name, processor.getType(attribute.SimpleType.Restriction.Base))
				} else if attribute.SimpleType.List != nil {
					processor.typeBuilder.AddAttribute(structName, *attribute.Name, processor.getType(attribute.SimpleType.List.ItemType))
				} else {
					// TODO
					println("Union inside attribute is not supported")
				}
			} else if attribute.Type != nil {
				processor.typeBuilder.AddAttribute(structName, *attribute.Name, processor.getType(*attribute.Type))
			}
		} else if attribute.Ref != nil {
			processor.typeBuilder.AddAttribute(structName, *attribute.Ref, processor.getType(*attribute.Ref))
		} else {
			println("Attribute is skipped")
		}
	}
}

func (processor xsdProcessor) processSimpleContent(structName string, simpleContent elements.SimpleContent) {
	if simpleContent.Restriction != nil {
		processor.typeBuilder.AddCharDataField(structName, processor.getType(simpleContent.Restriction.Base))
	} else {
		processor.typeBuilder.AddCharDataField(structName, processor.getType(simpleContent.Extension.Base))
		processor.processAttributes(structName, simpleContent.Extension.Attributes)
		// TODO
		// attr group
	}
}

func (processor xsdProcessor) processComplexContent(structName string, complexContent elements.ComplexContent) {
	if complexContent.Extension != nil {
		processor.typeBuilder.AddParent(structName, complexContent.Extension.Base)
		if complexContent.Extension.Group != nil {
			processor.processGroup(structName, *complexContent.Extension.Group)
		} else if complexContent.Extension.All != nil {
			processor.processAll(structName, *complexContent.Extension.All)
		} else if complexContent.Extension.Choice != nil {
			processor.processChoice(structName, *complexContent.Extension.Choice)
		} else if complexContent.Extension.Sequence != nil {
			processor.processSequence(structName, *complexContent.Extension.Sequence)
		}

		processor.processAttributes(structName, complexContent.Extension.Attributes)
	} else {
		processor.typeBuilder.AddParent(structName, complexContent.Restriction.Base)
		if complexContent.Restriction.Group != nil {
			processor.processGroup(structName, *complexContent.Restriction.Group)
		} else if complexContent.Restriction.All != nil {
			processor.processAll(structName, *complexContent.Restriction.All)
		} else if complexContent.Restriction.Choice != nil {
			processor.processChoice(structName, *complexContent.Restriction.Choice)
		} else if complexContent.Restriction.Sequence != nil {
			processor.processSequence(structName, *complexContent.Restriction.Sequence)
		}

		processor.processAttributes(structName, complexContent.Restriction.Attributes)
	}
}

func (processor xsdProcessor) processComplexType(structName *string, complexType elements.ComplexType) {
	var name string
	if structName != nil {
		name = *structName
	} else if complexType.Name != nil {
		name = *complexType.Name
	} else {
		return
	}
	processor.typeBuilder.AddType(name, "struct", "")
	if complexType.SimpleContent != nil {
		processor.processSimpleContent(name, *complexType.SimpleContent)
	} else if complexType.ComplexContent != nil {
		processor.processComplexContent(name, *complexType.ComplexContent)
	} else if complexType.Group != nil {
		processor.processGroup(name, *complexType.Group)
	} else if complexType.All != nil {
		processor.processAll(name, *complexType.All)
	} else if complexType.Choice != nil {
		processor.processChoice(name, *complexType.Choice)
	} else if complexType.Sequence != nil {
		processor.processSequence(name, *complexType.Sequence)
	} else {
		//fmt.Printf("Element is skipped: %s\n", *element.Name)
	}

	// attr
	processor.processAttributes(name, complexType.Attributes)
	// TODO
	// attr group
}

func (processor xsdProcessor) isArrayElement(element elements.Element) bool {
	return element.MaxOccurs != nil
}

func (processor xsdProcessor) processElements(structName *string, elements []elements.Element) {
	for _, element := range elements {
		if element.Name != nil {
			if element.Type == nil {
				if element.SimpleType != nil {
					processor.processSimpleType(structName, element.Name, *element.SimpleType)
				} else if element.ComplexType != nil {
					if structName != nil {
						processor.typeBuilder.AddField(*structName, *element.Name, *element.Name, processor.isArrayElement(element))
					}
					processor.processComplexType(element.Name, *element.ComplexType)
				} else {
					fmt.Printf("Element have doesn't have type, setting as string: %s\n", *element.Name)
					if structName == nil {
						processor.typeBuilder.AddType(*element.Name, "type", "string")
					} else {
						processor.typeBuilder.AddField(*structName, *element.Name, "string", processor.isArrayElement(element))
					}
				}
			} else {
				if structName == nil {
					processor.typeBuilder.AddType(*element.Name, "type", processor.getType(*element.Type))
				} else {
					processor.typeBuilder.AddField(*structName, *element.Name, processor.getType(*element.Type), processor.isArrayElement(element))
				}
			}
		} else if element.Ref != nil {
			if structName == nil {
				processor.typeBuilder.AddType(*element.Ref, "type", processor.getType(*element.Ref))
			} else {
				processor.typeBuilder.AddField(*structName, *element.Ref, processor.getType(*element.Ref), processor.isArrayElement(element))
			}
		} else {
			println("element skipped")
		}
	}
}

func (processor xsdProcessor) getType(inputType string) string {
	typeMap := map[string][]string{
		"string": {
			"string",
			"duration",
			"dateTime",
			"date",
			"time",
			"anyType",
			"anySimpleType",
			"gYearMonth",
			"gYear",
			"gMonthDay",
			"gDay",
			"gMonth",
			"base64Binary",
			"hexBinary",
			"anyURI",
			"normalizedString",
			"token",
			"NMTOKEN",
			"Name",
			"language",
			"NCName",
			"ENTITY",
			"IDREF",
			"ID",
			"NMTOKENS",
			"ENTITIES",
			"IDREFS",
			"QName",
			"NOTATION",
		},
		"bool": {
			"boolean",
		},
		"int": {
			"integer",
			"negativeInteger",
			"nonNegativeInteger",
			"nonPositiveInteger",
			"positiveInteger",
		},
		"int8": {
			"byte",
		},
		"int16": {
			"short",
		},
		"int32": {
			"int",
		},
		"int64": {
			"long",
		},
		"uint8": {
			"unsignedByte",
		},
		"uint16": {
			"unsignedShort",
		},
		"uint32": {
			"unsignedInt",
		},
		"uint64": {
			"unsignedLong",
		},
		"float64": {
			"decimal",
			"double",
			"float",
		},
	}
	for goType, types := range typeMap {
		for _, xmlType := range types {
			if processor.Prefix+xmlType == inputType {
				return goType
			}
		}
	}
	return inputType
}
