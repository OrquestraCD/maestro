package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestCreateTagsPlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Copy Image","mainSteps":[{"action":"aws:createTags","name":"createTags","inputs":{"ResourceIds":["ami-23l4k32l"],"Tags":[{"Key":"foo","Value":"bar"}]}}]}`

	input := CreateTagsPluginInput{
		ResourceIDs: []string{"ami-23l4k32l"},
		Tags: []Tag{{
			Key:   "foo",
			Value: "bar",
		}},
	}
	plugin := Plugin{
		Action: CreateTagsPluginAction,
		Name:   "createTags",
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
