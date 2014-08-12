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

type testCase func(t *testing.T, definition Definition)

func (test testCase) call(t *testing.T, definition Definition) {
	test(t, definition)
}

func TestUnmarshal(t *testing.T) {
	var definition Definition = testUnmarshalFromFile(t, "./testdata/wsdl.xml")

	var tests = []testCase{
		testCase(testUnmarshalDocumentation),
		testCase(testUnmarshalMessage),
		testCase(testUnmarshalPartInMessage),
		testCase(testUnmarshalPortType),
		testCase(testUnmarshalOperationInPortType),
		testCase(testUnmarshalOperationInputInPortType),
		testCase(testUnmarshalOperationOutputInPortType),
		testCase(testUnmarshalOperationFault),
		testCase(testUnmarshalSchema),
	}

	for _, test := range tests {
		test.call(t, definition)
	}
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
	var operations []WSDLOperation = definition.PortType.Operations

	if operations[0].Name != "ReadRetlWS" {
		t.Errorf("expect \"ReadRetlWS\" but was %s", operations[0].Name)
	}
}

func testUnmarshalOperationInputInPortType(t *testing.T, definition Definition) {
	var inputOperation InputOperation = definition.PortType.Operations[0].Input

	if inputOperation.Message != "tns:ReadRetlWSInput" {
		t.Errorf("expect \"tns:ReadRetlWSInput\" but was %s", inputOperation.Message)
	}
}

func testUnmarshalOperationOutputInPortType(t *testing.T, definition Definition) {
	var outputOperation OutputOperation = definition.PortType.Operations[0].Output

	if outputOperation.Message != "tns:ReadRetlWSOutput" {
		t.Errorf("expect \"tns:ReadRetlWSOutput\" but was %s", outputOperation.Message)
	}
}

func testUnmarshalOperationFault(t *testing.T, definition Definition) {
	var faultOperation FaultOperation = definition.PortType.Operations[0].Fault

	if faultOperation.Message != "tns:ReadRetlWSError" {
		t.Errorf("expect \"tns:ReadRetlWSError\" but was %s", faultOperation.Message)
	}

	if faultOperation.Name != "ReadRetlWSError" {
		t.Errorf("expect \"ReadRetlWSError\" but was %s", faultOperation.Name)
	}
}

func testUnmarshalSchema(t *testing.T, definition Definition) {
	var schema Schema = definition.Types.Schema

	if schema.AttributeFormDefault != "unqualified" {
		t.Errorf("expect \"unqualified\" but was %s", schema.AttributeFormDefault)
	}

	if schema.ElementFormDefault != "qualified" {
		t.Errorf("expect \"unqualified\" but was %s", schema.ElementFormDefault)
	}

	if schema.TargetNamespace != "urn:pack.INReadRetlWS_typedef.salt11" {
		t.Errorf("expect \"urn:pack.INReadRetlWS_PortType.salt11\" but was %s", schema.TargetNamespace)
	}
}
