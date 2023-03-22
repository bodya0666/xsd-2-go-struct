package elements

import "encoding/xml"

// All https://www.w3schools.com/xml/el_all.asp
type All struct {
	Id         *string     `xml:"id,attr"`
	MaxOccurs  *string     `xml:"maxOccurs,attr"`
	MinOccurs  *string     `xml:"minOccurs,attr"`
	Annotation *Annotation `xml:"annotation"`
	Elements   []Element   `xml:"element"`
}

// Annotation https://www.w3schools.com/xml/el_annotation.asp
type Annotation struct {
	Id             *string         `xml:"id,attr"`
	Appinfo        []Appinfo       `xml:"appinfo"`
	Documentations []Documentation `xml:"documentation"`
}

// Any https://www.w3schools.com/xml/el_any.asp
type Any struct {
	Id              *string     `xml:"id,attr"`
	MaxOccurs       *string     `xml:"maxOccurs,attr"`
	MinOccurs       *string     `xml:"minOccurs,attr"`
	Namespace       *string     `xml:"namespace,attr"`
	ProcessContents *string     `xml:"processContents,attr"`
	Annotation      *Annotation `xml:"annotation"`
}

// AnyAttribute https://www.w3schools.com/xml/el_anyattribute.asp
type AnyAttribute struct {
	Id              *string     `xml:"id,attr"`
	Namespace       *string     `xml:"namespace,attr"`
	ProcessContents *string     `xml:"processContents,attr"`
	Annotation      *Annotation `xml:"annotation"`
}

// Appinfo https://www.w3schools.com/xml/el_appinfo.asp
type Appinfo struct {
	Source *string `xml:"source,attr"`
	Value  string  `xml:",chardata"`
}

// Attribute https://www.w3schools.com/xml/el_attribute.asp
type Attribute struct {
	Default    *string     `xml:"default,attr"`
	Fixed      *string     `xml:"fixed,attr"`
	Form       *string     `xml:"form,attr"`
	Id         *string     `xml:"id,attr"`
	Name       *string     `xml:"name,attr"`
	Ref        *string     `xml:"ref,attr"`
	Type       *string     `xml:"type,attr"`
	Use        *string     `xml:"use,attr"`
	Annotation *Annotation `xml:"annotation"`
	SimpleType *SimpleType `xml:"simpleType"`
}

// AttributeGroup https://www.w3schools.com/xml/el_attributegroup.asp
type AttributeGroup struct {
	Id              *string          `xml:"id,attr"`
	Name            *string          `xml:"name,attr"`
	Ref             *string          `xml:"ref,attr"`
	Annotation      *Annotation      `xml:"annotation"`
	Attribute       []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
}

// Choice https://www.w3schools.com/xml/el_choice.asp
type Choice struct {
	Id         *string     `xml:"id,attr"`
	MaxOccurs  *string     `xml:"maxOccurs,attr"`
	MinOccurs  *string     `xml:"minOccurs,attr"`
	Annotation *Annotation `xml:"annotation"`
	Elements   []Element   `xml:"element"`
	Groups     []Group     `xml:"group"`
	Choices    []Choice    `xml:"choice"`
	Sequences  []Sequence  `xml:"sequence"`
	Any        []Any       `xml:"any"`
}

// ComplexContent https://www.w3schools.com/xml/el_complexcontent.asp
type ComplexContent struct {
	Id          *string                 `xml:"id,attr"`
	Mixed       bool                    `xml:"mixed,attr"`
	Annotation  *Annotation             `xml:"annotation"`
	Restriction *ComplexTypeRestriction `xml:"restriction"`
	Extension   *Extension              `xml:"extension"`
}

// ComplexType https://www.w3schools.com/xml/el_complextype.asp
type ComplexType struct {
	Id              *string          `xml:"id,attr"`
	Name            *string          `xml:"name,attr"`
	Abstract        *bool            `xml:"abstract,attr"`
	Mixed           *bool            `xml:"mixed,attr"`
	Block           *string          `xml:"block,attr"`
	Final           *string          `xml:"final,attr"`
	Annotation      *Annotation      `xml:"annotation"`
	SimpleContent   *SimpleContent   `xml:"simpleContent"`
	ComplexContent  *ComplexContent  `xml:"complexContent"`
	Group           *Group           `xml:"group"`
	All             *All             `xml:"all"`
	Choice          *Choice          `xml:"choice"`
	Sequence        *Sequence        `xml:"sequence"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
}

// Documentation https://www.w3schools.com/xml/el_documentation.asp
type Documentation struct {
	Source  *string `xml:"source,attr"`
	XmlLang *string `xml:"xml:lang,attr"`
	Value   string  `xml:",chardata"`
}

// Element https://www.w3schools.com/xml/el_element.asp
type Element struct {
	Id                *string      `xml:"id,attr"`
	Name              *string      `xml:"name,attr"`
	Ref               *string      `xml:"ref,attr"`
	Type              *string      `xml:"type,attr"`
	SubstitutionGroup *string      `xml:"substitutionGroup,attr"`
	Default           *string      `xml:"default,attr"`
	Fixed             *string      `xml:"fixed,attr"`
	Form              *string      `xml:"form,attr"`
	MaxOccurs         *string      `xml:"maxOccurs,attr"`
	MinOccurs         *string      `xml:"minOccurs,attr"`
	Nillable          *bool        `xml:"nillable,attr"`
	Abstract          *bool        `xml:"abstract,attr"`
	Block             *string      `xml:"block,attr"`
	Final             *string      `xml:"final,attr"`
	Annotation        *Annotation  `xml:"annotation"`
	ComplexType       *ComplexType `xml:"complexType"`
	SimpleType        *SimpleType  `xml:"simpleType"`
	Unique            []Unique     `xml:"unique"`
	Keys              []Key        `xml:"key"`
	Keyref            []Keyref     `xml:"keyref"`
}

// Extension https://www.w3schools.com/xml/el_extension.asp
type Extension struct {
	Id              *string          `xml:"id,attr"`
	Base            string           `xml:"base,attr"`
	Annotation      *Annotation      `xml:"annotation"`
	Group           *Group           `xml:"group"`
	All             *All             `xml:"all"`
	Choice          *Choice          `xml:"choice"`
	Sequence        *Sequence        `xml:"sequence"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
}

// Field https://www.w3schools.com/xml/el_field.asp
type Field struct {
	Id         *string     `xml:"id,attr"`
	XPath      string      `xml:"xpath,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// Group https://www.w3schools.com/xml/el_group.asp
type Group struct {
	Id         *string     `xml:"id,attr"`
	Name       *string     `xml:"name,attr"`
	Ref        *string     `xml:"ref,attr"`
	MaxOccurs  *string     `xml:"maxOccurs,attr"`
	MinOccurs  *string     `xml:"minOccurs,attr"`
	Annotation *Annotation `xml:"annotation"`
	All        *All        `xml:"all"`
	Choice     *Choice     `xml:"choice"`
	Sequence   *Sequence   `xml:"sequence"`
}

// Import https://www.w3schools.com/xml/el_import.asp
type Import struct {
	Id             *string     `xml:"id,attr"`
	Namespace      *string     `xml:"namespace,attr"`
	SchemaLocation *string     `xml:"schemaLocation,attr"`
	Annotation     *Annotation `xml:"annotation"`
}

// Include https://www.w3schools.com/xml/el_include.asp
type Include struct {
	XMLName        xml.Name
	ID             *string `xml:"ID,attr"`
	SchemaLocation string  `xml:"schemaLocation,attr"`
}

// Key https://www.w3schools.com/xml/el_key.asp
type Key struct {
	Id         *string     `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	Annotation *Annotation `xml:"annotation"`
	Fields     []Field     `xml:"field"`
	Selectors  []Selector  `xml:"selector"`
}

// Keyref https://www.w3schools.com/xml/el_keyref.asp
type Keyref struct {
	Id         *string     `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	Refer      string      `xml:"refer,attr"`
	Annotation *Annotation `xml:"annotation"`
	Fields     []Field     `xml:"field"`
	Selectors  []Selector  `xml:"selector"`
}

// List https://www.w3schools.com/xml/el_list.asp
type List struct {
	Id         *string     `xml:"id,attr"`
	ItemType   string      `xml:"itemType"`
	Annotation *Annotation `xml:"annotation"`
	SimpleType *SimpleType `xml:"simpleType"`
}

// Notation https://www.w3schools.com/xml/el_notation.asp
type Notation struct {
	Id     *string `xml:"id,attr"`
	Name   string  `xml:"name,attr"`
	Public string  `xml:"public,attr"`
	System string  `xml:"system,attr"`
}

// Redefine https://www.w3schools.com/xml/el_redefine.asp
type Redefine struct {
	Id              *string          `xml:"id,attr"`
	SchemaLocation  string           `xml:"schemaLocation,attr"`
	Annotations     []Annotation     `xml:"annotation"`
	SimpleTypes     []SimpleType     `xml:"simpleType"`
	ComplexTypes    []ComplexType    `xml:"complexTypes"`
	Groups          []Group          `xml:"group"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
}

// ComplexTypeRestriction https://www.w3schools.com/xml/el_restriction.asp
type ComplexTypeRestriction struct {
	Id              *string          `xml:"id,attr"`
	Base            string           `xml:"base,attr"`
	Annotation      *Annotation      `xml:"annotation"`
	Group           *Group           `xml:"group"`
	All             *All             `xml:"all"`
	Choice          *Choice          `xml:"choice"`
	Sequence        *Sequence        `xml:"sequence"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
}

// SimpleTypeRestriction https://www.w3schools.com/xml/el_restriction.asp
type SimpleTypeRestriction struct {
	Id             *string               `xml:"id,attr"`
	Base           string                `xml:"base,attr"`
	Enumeration    []RestrictionDataType `xml:"enumeration"`
	FractionDigits []RestrictionDataType `xml:"fractionDigits"`
	Length         []RestrictionDataType `xml:"length"`
	MaxExclusive   []RestrictionDataType `xml:"maxExclusive"`
	MaxInclusive   []RestrictionDataType `xml:"maxInclusive"`
	MaxLength      []RestrictionDataType `xml:"maxLength"`
	MinExclusive   []RestrictionDataType `xml:"minExclusive"`
	MinInclusive   []RestrictionDataType `xml:"minInclusive"`
	MinLength      []RestrictionDataType `xml:"minLength"`
	Pattern        []RestrictionDataType `xml:"pattern"`
	TotalDigits    []RestrictionDataType `xml:"totalDigits"`
	WhiteSpace     []RestrictionDataType `xml:"whiteSpace"`
}

// SimpleContentRestriction https://www.w3schools.com/xml/el_restriction.asp
type SimpleContentRestriction struct {
	Id              *string               `xml:"id,attr"`
	Base            string                `xml:"base,attr"`
	Enumeration     []RestrictionDataType `xml:"enumeration"`
	FractionDigits  []RestrictionDataType `xml:"fractionDigits"`
	Length          []RestrictionDataType `xml:"length"`
	MaxExclusive    []RestrictionDataType `xml:"maxExclusive"`
	MaxInclusive    []RestrictionDataType `xml:"maxInclusive"`
	MaxLength       []RestrictionDataType `xml:"maxLength"`
	MinExclusive    []RestrictionDataType `xml:"minExclusive"`
	MinInclusive    []RestrictionDataType `xml:"minInclusive"`
	MinLength       []RestrictionDataType `xml:"minLength"`
	Pattern         []RestrictionDataType `xml:"pattern"`
	TotalDigits     []RestrictionDataType `xml:"totalDigits"`
	WhiteSpace      []RestrictionDataType `xml:"whiteSpace"`
	Attributes      []Attribute           `xml:"attribute"`
	AttributeGroups []AttributeGroup      `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute         `xml:"anyAttribute"`
}

// RestrictionDataType https://www.w3schools.com/xml/schema_facets.asp
type RestrictionDataType struct {
	Value string `xml:"base,attr"`
}

// Schema https://www.w3schools.com/xml/el_schema.asp
type Schema struct {
	Id                   *string          `xml:"id,attr"`
	AttributeFormDefault *string          `xml:"attributeFormDefault,attr"`
	ElementFormDefault   *string          `xml:"elementFormDefault,attr"`
	BlockDefault         *string          `xml:"blockDefault,attr"`
	FinalDefault         *string          `xml:"finalDefault,attr"`
	TargetNamespace      *string          `xml:"targetNamespace,attr"`
	Version              *string          `xml:"version,attr"`
	Xmlns                string           `xml:"xmlns,attr"`
	Includes             []Include        `xml:"include"`
	Imports              []Import         `xml:"import"`
	Redefines            []Redefine       `xml:"redefine"`
	SimpleTypes          []SimpleType     `xml:"simpleType"`
	ComplexTypes         []ComplexType    `xml:"complexType"`
	Groups               []Group          `xml:"group"`
	Attributes           []Attribute      `xml:"attribute"`
	AttributeGroups      []AttributeGroup `xml:"attributeGroup"`
	Notations            []Notation       `xml:"notation"`
	Annotations          []Annotation     `xml:"annotation"`
	Elements             []Element        `xml:"element"`
}

// Selector https://www.w3schools.com/xml/el_selector.asp
type Selector struct {
	Id         *string     `xml:"id,attr"`
	XPath      string      `xml:"xpath,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// Sequence https://www.w3schools.com/xml/el_sequence.asp
type Sequence struct {
	Id         *string     `xml:"id,attr"`
	MaxOccurs  *string     `xml:"maxOccurs,attr"`
	MinOccurs  *string     `xml:"minOccurs,attr"`
	Annotation *Annotation `xml:"annotation"`
	Elements   []Element   `xml:"element"`
	Groups     []Group     `xml:"group"`
	Choices    []Choice    `xml:"choice"`
	Sequences  []Sequence  `xml:"sequence"`
	Any        []Any       `xml:"any"`
}

// SimpleContent https://www.w3schools.com/xml/el_simpleContent.asp
type SimpleContent struct {
	Id          *string                   `xml:"id,attr"`
	Annotation  *Annotation               `xml:"annotation"`
	Restriction *SimpleContentRestriction `xml:"restriction"`
	Extension   *Extension                `xml:"extension"`
}

// SimpleType https://www.w3schools.com/xml/el_simpletype.asp
type SimpleType struct {
	Id          *string                `xml:"id,attr"`
	Name        *string                `xml:"name,attr"`
	Annotation  *Annotation            `xml:"annotation"`
	Restriction *SimpleTypeRestriction `xml:"restriction"`
	List        *List                  `xml:"list"`
	Union       *Union                 `xml:"union"`
}

// Union https://www.w3schools.com/xml/el_union.asp
type Union struct {
	Id          *string      `xml:"id,attr"`
	MemberTypes *string      `xml:"memberTypes,attr"`
	Annotation  *Annotation  `xml:"annotation"`
	SimpleTypes []SimpleType `xml:"simpleType"`
}

// Unique https://www.w3schools.com/xml/el_unique.asp
type Unique struct {
	Id         *string     `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	Annotation *Annotation `xml:"annotation"`
	Fields     []Field     `xml:"field"`
	Selectors  []Selector  `xml:"selector"`
}
