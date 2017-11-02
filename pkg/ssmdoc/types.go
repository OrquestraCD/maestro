package ssmdoc

import (
	"encoding/json"
)

type List struct {
	reference string
	value     []string
}

func ListReference(reference string) List {
	return List{reference: "{{ " + reference + " }}"}
}

func ListValue(value []string) List {
	return List{value: value}
}

func (s List) MarshalJSON() ([]byte, error) {
	if s.reference != "" {
		return json.Marshal(s.reference)
	}

	return json.Marshal(s.value)
}
