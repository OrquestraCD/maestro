package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestRunInstancesPlugin(t *testing.T) {
	want := `{"assumeRole":"fauxRole","schemaVersion":"0.3","description":"Test Run Instances","mainSteps":[{"action":"aws:runInstances","name":"runInstances","inputs":{"ImageId":"ami-12345678","IamInstanceProfileName":"MyTestInstanceProfile","InstanceType":"t2.medium","MaxInstanceCount":1,"MinInstanceCount":1}}]}`

	input := RunInstancesPluginInput{
		ImageID:                "ami-12345678",
		InstanceType:           "t2.medium",
		MinInstanceCount:       1,
		MaxInstanceCount:       1,
		IamInstanceProfileName: "MyTestInstanceProfile",
	}
	plugin := Plugin{
		Action: RunInstancesPluginAction,
		Name:   "runInstances",
		Inputs: input,
	}

	doc, err := NewAutomationDocument("Test Run Instances", "fauxRole")
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
