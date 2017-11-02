package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestApplicationsPlugin(t *testing.T) {
	want := `{"schemaVersion":"2.0","description":"Test Applications","mainSteps":[{"action":"aws:applications","name":"applicationsPlugin","inputs":{"action":"Install","source":"file:///package.msi"}}]}`

	input := AWSApplicationsPluginInput{
		Action: "Install",
		Source: "file:///package.msi",
	}
	applications := Plugin{
		Action: AWSApplicationsPluginAction,
		Name:   "applicationsPlugin",
		Inputs: input,
	}

	doc, err := NewDocument("Test Applications")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	doc.AddStep(applications)

	resp, err := json.Marshal(doc)
	if err != nil {
		t.Error(err.Error())
	}

	if want != string(resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"\n", want, resp)
	}
}
