package ssmdoc

const (
	CreateTagsPluginAction = "aws:createTags"
)

type CreateTagsPluginInput struct {
	ResourceType string   `json:"ResourceType,omitempty"`
	ResourceIDs  []string `json:"ResourceIds" required:"true"`
	Tags         []Tag    `required:"true"`
}

type Tag struct {
	Key   string `json:"Key" required:"true"`
	Value string `json:"Value" required:"true"`
}
