package ssmdoc

const (
	CopyImagePluginAction = "aws:copyImage"
)

type CopyImagePluginInput struct {
	ClientToken      string `json:"ClientToken,omitempty"`
	Encrypted        bool   `json:"Encrypted,omitempty"`
	KMSKeyID         string `json:"KmsKeyId,omitempty"`
	ImageDescription string `json:"ImageDescription,omitempty"`
	ImageName        string `json:"ImageName" required:"true"`
	SourceImageID    string `json:"SourceImageId" required:"true"`
	SourceRegion     string `json:"SourceRegion" required:"true"`
}
