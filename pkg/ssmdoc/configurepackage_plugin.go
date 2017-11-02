package ssmdoc

const (
	AWSConfigurePackagePluginAction = "aws:configurePackage"
)

type AWSConfigurePackagePluginInput struct {
	Name    string `json:"name" required:"true"`
	Action  string `json:"action" required:"true"`
	Version string `json:"version,omitempty"`
}
