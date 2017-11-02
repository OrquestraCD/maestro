package fileutil

import "testing"

func TestReadFileToString(t *testing.T) {
	testFile := "bom.ps1"
	testContent := "This file has a BOM.\n"

	result, err := ReadFileToString(testFile)
	if err != nil {
		t.Errorf("Unexpected error %v\n", err)
	}

	if result != testContent {
		t.Errorf("Unexpected result returned: \"%s\"", result)
	}
}
