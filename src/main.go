package main

import (
	"main/xsdProcessor"
)

func main() {
	xsdProcessor.
		Create("xsd:", "types").
		Process("/www/xsd-struct/amz-product-xsd/MechanicalFasteners.xsd")
}
