package ssmdoc

const (
	RunCommandPluginAction = "aws:runCommand"
)

type RunCommandPluginInput struct {
	Comment            string                 `json:"Comment,omitempty"`
	DocumentName       string                 `json:"DocumentName" required:"true"`
	DocumentHash       string                 `json:"DocumentHash,omitempty"`
	DocumentHashType   string                 `json:"DocumentHashType,omitempty"`
	InstanceIDs        []string               `json:"InstanceIds,omitempty"`
	OutputS3BucketName string                 `json:"OutputS3BucketName,omitempty"`
	OutputS3KeyPrefix  string                 `json:"OutputS3KeyPrefix,omitempty"`
	Parameters         map[string]interface{} `json:"Parameters,omitempty"`
	ServiceRoleArn     string                 `json:"ServiceRoleArn,omitempty"`
	TimeoutSeconds     int                    `json:"TimeoutSeconds,omitempty"`
}
