package builder

type TypeBuilder interface {
	AddType(structName, itemType, itemTypeName string)
	AddField(structName string, fieldName, fieldType string, isArray bool)
	AddCharDataField(structName string, fieldType string)
	AddParent(structName string, parentStructName string)
	AddAttribute(structName string, fieldName, fieldType string)
	Build() string
}

type Type struct {
	Name     string
	Type     string
	TypeName string
	Fields   []string
}
