package wsdl

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func testRead(t *testing.T, filename string) []byte {
	var b, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return []byte{}
	}
	return b
}

func testUnmarshal(t *testing.T, b []byte) Definition {
	definition, err := Unmarshal(b)
	if err != nil {
		t.Error(err)
		return Definition{}
	}
	return definition
}

func testUnmarshalFromFile(t *testing.T, filename string) (definition Definition) {
	b := testRead(t, "./testdata/wsdl.xml")
	definition = testUnmarshal(t, b)
	return
}

func TestUnmarshal(t *testing.T) {
	var definition Definition = testUnmarshalFromFile(t, "./testdata/wsdl.xml")

	testUnmarshalDocumentation(t, definition)
	testUnmarshalMessage(t, definition)
	testUnmarshalPartInMessage(t, definition)
	testUnmarshalPortType(t, definition)
	testUnmarshalOperationInPortType(t, definition)
	testUnmarshalOperationInputInPortType(t, definition)
}

func testUnmarshalDocumentation(t *testing.T, definition Definition) {
	if definition.Documentation != "Generated at 08-11-2014 16:33:06:253" {
		t.Errorf("expect \"Generated at 08-11-2014 16:33:06:253\" but was %s", definition.Documentation)
	}
}

func testUnmarshalMessage(t *testing.T, definition Definition) {
	messages := definition.Messages

	if messages[0].Name != "ReadRetlWSInput" {
		t.Errorf("expect \"ReadRetlWSInput\" but was %s", messages[0].Name)
	}
	if messages[1].Name != "ReadRetlWSOutput" {
		t.Errorf("expect \"ReadRetlWSOutput\" but was %s", messages[0].Name)
	}
	if messages[2].Name != "ReadRetlWSError" {
		t.Errorf("expect \"ReadRetlWSError\" but was %s", messages[0].Name)
	}
}

func testUnmarshalPartInMessage(t *testing.T, definition Definition) {
	messages := definition.Messages
	parts := []Part{
		messages[0].Part,
		messages[1].Part,
		messages[2].Part,
	}

	expectedParts := []Part{
		Part{
			Element: "tuxtype:ReadRetlWS",
			Name:    "FML32",
		},
		Part{
			Element: "tuxtype:ReadRetlWSResponse",
			Name:    "FML32",
		},
		Part{
			Element: "tuxtype:ReadRetlWSFault",
			Name:    "FML32",
		},
	}

	if !reflect.DeepEqual(parts, expectedParts) {
		t.Errorf("expect parts %v but was %v", expectedParts, parts)
	}
}

func testUnmarshalPortType(t *testing.T, definition Definition) {
	portType := definition.PortType

	if portType.Name != "INReadRetlWS_PortType" {
		t.Errorf("expect \"INReadRetlWS_PortType\" but was %s", portType.Name)
	}
}

func testUnmarshalOperationInPortType(t *testing.T, definition Definition) {
	var operation WSDLOperation = definition.PortType.Operation

	if operation.Name != "ReadRetlWS" {
		t.Errorf("expect \"ReadRetlWS\" but was %s", operation.Name)
	}
}

func testUnmarshalOperationInputInPortType(t *testing.T, definition Definition) {
	var inputOperation InputOperation = definition.PortType.Operation.Input

	if inputOperation.Message != "tns:ReadRetlWSInput" {
		t.Errorf("expect \"tns:ReadRetlWSInput\" but was %s", inputOperation.Message)
	}
}
