package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestConfigurePackagePlugin(t *testing.T) {
	want := `{"schemaVersion":"2.0","description":"Test Configure Package","mainSteps":[{"action":"aws:configurePackage","name":"installApache","inputs":{"name":"apache2","action":"Install"}}]}`

	inputs := AWSConfigurePackagePluginInput{
		Action: "Install",
		Name:   "apache2",
	}

	configurePackage := Plugin{
		Action: AWSConfigurePackagePluginAction,
		Name:   "installApache",
		Inputs: inputs,
	}

	doc, err := NewDocument("Test Configure Package")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	doc.AddStep(configurePackage)

	resp, err := json.Marshal(doc)
	if err != nil {
		t.Error(err.Error())
	}

	if want != string(resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"\n", want, resp)
	}
}
