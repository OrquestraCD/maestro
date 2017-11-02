package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestDeleteImagePlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Copy Image","mainSteps":[{"action":"aws:createImage","name":"deleteImage","inputs":{"ImageId":"ami-lkj23423"}}]}`

	input := DeleteImagePluginInput{
		ImageID: "ami-lkj23423",
	}
	plugin := Plugin{
		Action: DeleteImagePluginAction,
		Name:   "deleteImage",
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
