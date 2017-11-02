package scripts

import (
	"reflect"
	"testing"
)

func TestEscapeString(t *testing.T) {
	testStr := `echo "hello world"`
	want := `echo \"hello world\"`

	got := EscapeString(testStr)
	if want != got {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, got)
	}
}

func TestScriptToSSMCommands(t *testing.T) {
	testScript := `echo "hello world"
echo "this is my second line"`

	want := []string{
		`echo "hello world"`,
		`echo "this is my second line"`,
	}

	got := ScriptToSSMCommands(testScript)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, got)
	}
}
