package ssmdoc

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestListReference(t *testing.T) {
	paramName := "command"
	want := []byte{34, 123, 123, 32, 99, 111, 109, 109, 97, 110, 100, 32, 125, 125, 34}

	list := ListReference(paramName)

	resp, err := json.Marshal(list)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	if !reflect.DeepEqual(want, resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, resp)
	}
}

func TestListValue(t *testing.T) {
	command := []string{"echo \"foo bar\""}
	want := []byte{91, 34, 101, 99, 104, 111, 32, 92, 34, 102, 111, 111, 32, 98, 97, 114, 92, 34, 34, 93}

	list := ListValue(command)

	resp, err := json.Marshal(list)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	if !reflect.DeepEqual(want, resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, resp)
	}
}
