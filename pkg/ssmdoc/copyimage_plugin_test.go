package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestCopyImagePlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Copy Image","mainSteps":[{"action":"aws:copyImage","name":"copyImage","inputs":{"ImageName":"Copy of my test image","SourceImageId":"ami-23l4k32l","SourceRegion":"us-east-2"}}]}`

	input := CopyImagePluginInput{
		SourceImageID: "ami-23l4k32l",
		SourceRegion:  "us-east-2",
		ImageName:     "Copy of my test image",
	}
	plugin := Plugin{
		Action: CopyImagePluginAction,
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
