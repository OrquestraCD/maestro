package json

import (
	"testing"
)

// Wrap JSON Marshal and unescape ampersand
func TestMarshal(t *testing.T) {
	expect := `{"SomeField":"test & and < and >"}`

	var testStruct = struct {
		SomeField string
	}{
		SomeField: "test & and < and >",
	}

	result, err := Marshal(&testStruct)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if string(result) != expect {
		t.Errorf("Expected \"%s\", Got \"%s\"", expect, result)
	}
}

func TestMarshalIndent(t *testing.T) {
	expect := `{
  "SomeField": "test & and < and >"
}`

	var testStruct = struct {
		SomeField string
	}{
		SomeField: "test & and < and >",
	}

	result, err := MarshalIndent(&testStruct, "", "  ")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if string(result) != expect {
		t.Errorf("Expected \"%s\", Got \"%s\"", expect, result)
	}
}
