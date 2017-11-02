package ssmdoc

const (
	AWSApplicationsPluginAction = "aws:applications"
)

type AWSApplicationsPluginInput struct {
	Action     string `json:"action" required:"true"`
	ID         string `json:"id,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Source     string `json:"source" required:"true"`
	SourceHash string `json:"sourceHash,omitempty"`
}
