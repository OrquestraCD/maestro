package ssmdoc

import (
	"encoding/json"
	"testing"
)

func TestUpdateSSMAgentPlugin(t *testing.T) {
	want := `{"schemaVersion":"2.0","description":"Test Update SSM Agent","mainSteps":[{"action":"aws:updateSsmAgent","name":"updateSSMAgent","inputs":{"agentName":"amazon-ssm-agent","source":"https://s3.us-east-1.amazonaws.com/aws-ssm-us-east-1/manifest.json"}}]}`

	input := AWSUpdateSSMAgentPluginInput{
		AgentName: "amazon-ssm-agent",
		Source:    "https://s3.us-east-1.amazonaws.com/aws-ssm-us-east-1/manifest.json",
	}

	updateSSMAgent := Plugin{
		Action: AWSUpdateSSMAgentPluginAction,
		Name:   "updateSSMAgent",
		Inputs: input,
	}

	doc, err := NewDocument("Test Update SSM Agent")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	doc.AddStep(updateSSMAgent)

	resp, err := json.Marshal(doc)
	if err != nil {
		t.Error(err.Error())
	}

	if want != string(resp) {
		t.Errorf("Expected \"%s\", Got \"%s\"\n", want, resp)
	}
}
