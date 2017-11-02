package ssmdoc

const (
	ChangeInstanceStatePluginAction = "aws:changeInstanceState"
)

type ChangeInstanceStatePluginInput struct {
	AdditionalInfo string   `json:"AdditionalInfo,omitempty"`
	InstanceIDs    []string `json:"InstanceIds" required:"true"`
	CheckStateOnly bool     `json:"CheckStateOnly,omitempty"`
	DesiredState   string   `json:"DesiredState" required:"true"`
	Force          bool     `json:"Force,omitempty"`
}
