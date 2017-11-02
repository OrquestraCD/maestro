package fileutil

import (
	"bytes"
	"io/ioutil"
)

// Read a file and return the contents as a string
func ReadFileToString(name string) (string, error) {
	fil, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	// Remove the BOM (Common for UTF-8 files saved in Windows)
	trimmedBytes := bytes.TrimLeft(fil, "\xef\xbb\xbf")

	return string(trimmedBytes), nil
}
