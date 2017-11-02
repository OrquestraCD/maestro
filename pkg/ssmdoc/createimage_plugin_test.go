package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestCreateImagePlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Copy Image","mainSteps":[{"action":"aws:createImage","name":"copyImage","inputs":{"InstanceId":"i-0123456789lkjl","ImageName":"AMI Created on{{global:DATE_TIME}}","NoReboot":true}}]}`

	input := CreateImagePluginInput{
		InstanceID: "i-0123456789lkjl",
		ImageName:  "AMI Created on{{global:DATE_TIME}}",
		NoReboot:   true,
	}
	plugin := Plugin{
		Action: CreateImagePluginAction,
		Name:   "copyImage",
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
