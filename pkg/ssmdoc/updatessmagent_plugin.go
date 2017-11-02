package ssmdoc

const (
	AWSUpdateSSMAgentPluginAction = "aws:updateSsmAgent"
)

type AWSUpdateSSMAgentPluginInput struct {
	AgentName      string `json:"agentName" required:"true"`
	Source         string `json:"source" required:"true"`
	AllowDowngrade string `json:"allowDowngrade,omitempty"`
	TargetVersion  string `json:"targetVersion,omitempty"`
}
