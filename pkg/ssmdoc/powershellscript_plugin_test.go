package ssmdoc

import "testing"

func TestRunPowerShellScriptPlugin(t *testing.T) {
	want := `{"schemaVersion":"2.0","description":"Test RunShell","mainSteps":[{"action":"aws:runPowerShellScript","name":"runPowerShellScript","inputs":{"runCommand":["dir"]}}]}`

	input := RunPowerShellScriptPluginInput{
		RunCommand: ListValue([]string{"dir"}),
	}

	plugin := Plugin{
		Action: RunPowerShellScriptPluginAction,
		Name:   "runPowerShellScript",
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
