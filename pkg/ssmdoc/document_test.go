package ssmdoc

import (
	"testing"
)

func TestSSMDocument(t *testing.T) {
	doc, err := NewDocument("Test document")
	if err != nil {
		t.Error(err.Error())
	}

	doc.AddParameter(Parameter{
		Type:     "StringList",
		MinItems: 1,
	})

	input := RunShellScriptPluginInput{
		RunCommand: ListReference("commands"),
	}
	runShell := Plugin{
		Action: RunShellScriptPluginAction,
		Name:   "runShellScript",
		Inputs: input,
	}

	err = doc.AddStep(runShell)
	if err != nil {
		t.Error(err.Error())
	}

	want := `{"schemaVersion":"2.0","description":"Test document","parameters":[{"type":"StringList","minItems":1}],"mainSteps":[{"action":"aws:runShellScript","name":"runShellScript","inputs":{"runCommand":"{{ commands }}"}}]}`

	result, err := doc.String()
	if err != nil {
		t.Error(err.Error())
	}

	if result != want {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, result)
	}
}

func TestAutomationDocument(t *testing.T) {
	doc, err := NewAutomationDocument("Test document", "testRoleArn")
	if err != nil {
		t.Error(err.Error())
	}

	doc.AddParameter(Parameter{
		Type:     "ImageId",
		MinItems: 1,
	})

	// Add Start Instance
	doc.AddStep(Plugin{
		Action: RunInstancesPluginAction,
		Name:   "runInstance",
		Inputs: RunInstancesPluginInput{
			ImageID:          "{{ImageId}}",
			MaxInstanceCount: 1,
			MinInstanceCount: 1,
		},
	})

	doc.AddStep(Plugin{
		Action: RunCommandPluginAction,
		Name:   "runCommand",
		Inputs: RunCommandPluginInput{
			DocumentName: "",
			Parameters: map[string]interface{}{
				"commands": ListValue([]string{"apt-get update -y && apt-get upgrade -y"}),
			},
		},
	})

	want := `{"assumeRole":"testRoleArn","schemaVersion":"0.3","description":"Test document","parameters":[{"type":"ImageId","minItems":1}],"mainSteps":[{"action":"aws:runInstances","name":"runInstance","inputs":{"ImageId":"{{ImageId}}","MaxInstanceCount":1,"MinInstanceCount":1}},{"action":"aws:runCommand","name":"runCommand","inputs":{"DocumentName":"","Parameters":{"commands":["apt-get update -y && apt-get upgrade -y"]}}}]}`
	got, err := doc.String()
	if err != nil {
		t.Error(err.Error())
	}

	if want != got {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, got)
	}
}
