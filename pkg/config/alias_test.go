package config

import "testing"

func TestDocument(t *testing.T) {
	alias := Alias{
		Description: "Various System Info",
		Command:     "top -n 1 %% df -h",
		Name:        "system_stuff",
		Type:        "bash",
	}

	result, err := alias.Document()
	if err != nil {
		t.Error(err.Error())
	}

	expect := `{"schemaVersion":"2.0","description":"Various System Info","mainSteps":[{"action":"aws:runShellScript","name":"runShellScript","inputs":{"runCommand":["top -n 1 %% df -h"]}}]}`
	if string(*result) != expect {
		t.Errorf("Expected: %s\n, Got %s\n", expect, *result)
	}
}

func TestValidateCommandFail(t *testing.T) {
	alias := Alias{
		Description: "Does things",
		Name:        "system_stuff",
	}

	if err := alias.validate(); err == nil {
		t.Errorf("expected error when no Command is set")
	}
}

func TestValidate(t *testing.T) {
	alias := Alias{
		Command:     "top -n 1 %% df -h",
		Description: "Gathers various system things.",
		Name:        "system_stuff",
		Type:        "bash",
	}

	if err := alias.validate(); err != nil {
		t.Error(err.Error())
	}
}
