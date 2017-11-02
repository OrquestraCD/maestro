package scripts

import "testing"

func TestDetectScriptByExtension_bash(t *testing.T) {
	testBash := "foo.sh"
	result := DetectScriptByExtension(testBash)

	if result != Bash {
		t.Errorf("Expected \"%s\" to be Bash", testBash)
	}
}

func TestDetectScriptByExtension_powershell(t *testing.T) {
	scriptName := "foo.ps1"
	result := DetectScriptByExtension(scriptName)

	if result != PowerShell {
		t.Errorf("Expected \"%s\" to be PowerShell", scriptName)
	}
}

func TestDetectScriptByExtension_unkown(t *testing.T) {
	scriptName := "foo"
	result := DetectScriptByExtension(scriptName)

	if result != Unknown {
		t.Errorf("Expected \"%s\" to be Unknown", scriptName)
	}
}
