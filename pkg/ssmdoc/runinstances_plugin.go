package ssmdoc

const (
	RunInstancesPluginAction = "aws:runInstances"
)

type RunInstancesPluginInput struct {
	AdditionalInfo                    string   `json:"AdditionalInfo,omitempty"`
	ClientToken                       string   `json:"ClientToken,omitempty"`
	DisableApiTermination             bool     `json:"DisableApiTermination,omitempty"`
	EbsOptimized                      bool     `json:"EbsOptimized,omitempty"`
	ImageID                           string   `json:"ImageId" required:"true"`
	IamInstanceProfileArn             string   `json:"IamInstanceProfileArn,omitempty"`
	IamInstanceProfileName            string   `json:"IamInstanceProfileName,omitempty"`
	InstanceInitiatedShutdownBehavior string   `json:"InstanceInitiatedShutdownBehavior,omitempty"`
	InstanceType                      string   `json:"InstanceType,omitempty"`
	KernelID                          string   `json:"KernelId,omitempty"`
	KeyName                           string   `json:"KeyName,omitempty"`
	MaxInstanceCount                  int      `json:"MaxInstanceCount,omitempty"`
	MinInstanceCount                  int      `json:"MinInstanceCount,omitempty"`
	Monitoring                        bool     `json:"Monitoring,omitempty"`
	SecurityGroupIDs                  []string `json:"SecurityGroupIds,omitempty"`
	SubnetID                          string   `json:"SubnetId,omitempty"`
	UserData                          string   `json:"UserData,omitempty"`
}
