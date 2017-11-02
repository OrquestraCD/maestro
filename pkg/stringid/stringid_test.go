package stringid

import "testing"

func TestGenerateID(t *testing.T) {
	if len(GenerateID()) != 16 {
		t.Error("expected length of 16")
	}
}
