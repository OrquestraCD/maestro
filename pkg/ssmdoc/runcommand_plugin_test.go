package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestRunCommandPlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Copy Image","mainSteps":[{"action":"aws:runCommand","name":"runCommand","inputs":{"DocumentName":"AWS-InstallPowerShellModule","Parameters":{"source":"https://my-s3-url.com/MyModule.zip","sourceHash":"ASDFWER12321WRW"}}}]}`

	input := RunCommandPluginInput{
		DocumentName: "AWS-InstallPowerShellModule",
		Parameters: map[string]interface{}{
			"source":     "https://my-s3-url.com/MyModule.zip",
			"sourceHash": "ASDFWER12321WRW",
		},
	}
	plugin := Plugin{
		Action: RunCommandPluginAction,
		Name:   "runCommand",
		Inputs: input,
	}

	doc, err := NewAutomationDocument("Test Copy Image", "fauxRole")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	doc.AddStep(plugin)

	resp, err := json.Marshal(doc)
	if err != nil {
		t.Error(err.Error())
	}

	if want != string(resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"\n", want, resp)
	}
}
