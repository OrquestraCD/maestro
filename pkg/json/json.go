package json

import (
	"bytes"
	"encoding/json"
)

const ampersandUnicode = "\\u0026"

var unicodeMap = map[string]string{
	"\\u0026": "&",
	"\\u003c": "<",
	"\\u003e": ">",
}

// Wrap JSON Marshal and unescape ampersand
func Marshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}

	return replaceUnicodeConversion(b), err
}

func MarshalIndent(v interface{}, prefix, index string) ([]byte, error) {
	b, err := json.MarshalIndent(v, prefix, index)
	if err != nil {
		return []byte{}, err
	}

	return replaceUnicodeConversion(b), err
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func replaceUnicodeConversion(data []byte) []byte {
	b := data

	for code, char := range unicodeMap {
		b = bytes.Replace(b, []byte(code), []byte(char), -1)
	}

	return b
}
