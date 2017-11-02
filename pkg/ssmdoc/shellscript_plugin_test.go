package ssmdoc

import "testing"

func TestShellScriptPlugin(t *testing.T) {
	want := `{"schemaVersion":"2.0","description":"Test RunShell","mainSteps":[{"action":"aws:runShellScript","name":"runShellScript","inputs":{"runCommand":["df -h"]}}]}`

	input := RunShellScriptPluginInput{
		RunCommand: ListValue([]string{"df -h"}),
	}

	plugin := Plugin{
		Action: RunShellScriptPluginAction,
		Name:   "runShellScript",
		Inputs: input,
	}

	doc, err := NewDocument("Test RunShell")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	if err := doc.AddStep(plugin); err != nil {
		t.Error(err.Error())
	}

	resp, err := doc.JSON()
	if err != nil {
		t.Error(err.Error())
	}

	if want != string(resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"\n", want, resp)
	}
}
