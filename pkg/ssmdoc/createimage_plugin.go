package ssmdoc

const (
	CreateImagePluginAction = "aws:createImage"
)

type CreateImagePluginInput struct {
	BlockDeviceMappings string `json:"BlockDeviceMappings,omitempty"`
	ImageDescription    string `json:"ImageDescription,omitempty"`
	InstanceID          string `json:"InstanceId" required:"true"`
	ImageName           string `json:"ImageName" required:"true"`
	NoReboot            bool   `json:"NoReboot,omitempty"`
}
