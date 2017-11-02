package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestChangeInstanceStatePlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Change Instance State","mainSteps":[{"action":"aws:changeInstanceState","name":"changeInstanceState","inputs":{"InstanceIds":["i-1234567890abcdef0"],"DesiredState":"stopped"}}]}`

	input := ChangeInstanceStatePluginInput{
		DesiredState: "stopped",
		InstanceIDs:  []string{"i-1234567890abcdef0"},
	}
	plugin := Plugin{
		Action: ChangeInstanceStatePluginAction,
		Name:   "changeInstanceState",
		Inputs: input,
	}

	doc, err := NewAutomationDocument("Test Change Instance State", "fauxRole")
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
