package ssmdoc

const (
	RunShellScriptPluginAction = "aws:runShellScript"
)

type RunShellScriptPluginInput struct {
	ID               string `json:"id,omitempty" required:"true"`
	RunCommand       List   `json:"runCommand"`
	WorkingDirectory string `json:"workingDirectory,omitempty" required:"true"`
	TimeoutSeconds   string `json:"timeoutSeconds,omitempty" required:"true"`
}
