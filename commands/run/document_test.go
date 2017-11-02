package run

import (
	"reflect"
	"testing"
)

func TestLoadParameters(t *testing.T) {
	input := "param1=value1 param2=value2"
	want := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	got, err := loadParameters(input, " ")
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected \"%+v\", Got \"%+v\"", want, got)
	}
}
