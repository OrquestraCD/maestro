package ssmdoc

const (
	DeleteImagePluginAction = "aws:createImage"
)

type DeleteImagePluginInput struct {
	ImageID string `json:"ImageId" required:"true"`
}
