package ssmdoc

import (
	"github.com/rackerlabs/maestro/pkg/json"
)

const (
	schemaVersion           = "2.0"
	automationSchemaVersion = "0.3"
)

type Parameter struct {
	Type           string   `json:"type"`
	Description    string   `json:"description,omitempty"`
	Default        string   `json:"default,omitempty"`
	AllowedValues  []string `json:"allowedValues,omitempty"`
	AllowedPattern string   `json:"allowedPattern,omitempty"`
	DisplayType    string   `json:"displayType,omitempty"`
	MinItems       int      `json:"minItems,omitempty"`
	MaxItems       int      `json:"maxItems,omitempty"`
	MinChars       int      `json:"minChars,omitempty"`
	MaxChars       int      `json:"maxChars,omitempty"`
}

type Plugin struct {
	Action         string      `json:"action" required:"true"`
	MaxAttempts    int         `json:"maxAttempts,omitempty"`
	Name           string      `json:"name" required:"true"`
	OnFailure      string      `json:"onFailure,omitempty"`
	Inputs         interface{} `json:"inputs" required:"true"`
	TimeoutSeconds int         `json:"timeoutSeconds,omitempty"`
}

type SSMDocument struct {
	AssumeRole    string      `json:"assumeRole,omitempty"` // Only used in schema 0.3
	SchemaVersion string      `json:"schemaVersion"`
	Description   string      `json:"description"`
	Parameters    []Parameter `json:"parameters,omitempty"`
	MainSteps     []Plugin    `json:"mainSteps,omitempty"`
}

func NewAutomationDocument(description, assumeRole string) (*SSMDocument, error) {
	doc := SSMDocument{
		AssumeRole:    assumeRole,
		SchemaVersion: automationSchemaVersion,
		Description:   description,
		MainSteps:     make([]Plugin, 0),
	}

	return &doc, nil
}

// Create a new SSM Document
func NewDocument(description string) (*SSMDocument, error) {
	doc := SSMDocument{
		SchemaVersion: schemaVersion,
		Description:   description,
		MainSteps:     make([]Plugin, 0),
	}

	return &doc, nil
}

// Add a step to a document
func (d *SSMDocument) AddStep(step Plugin) error {
	// Ensure that all required fields in a step are set
	if err := validateStep(step); err != nil {
		return err
	}

	d.MainSteps = append(d.MainSteps, step)
	return nil
}

func (d *SSMDocument) AddParameter(param Parameter) {
	d.Parameters = append(d.Parameters, param)
}

func (d *SSMDocument) String() (string, error) {
	resp, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (d *SSMDocument) JSON() ([]byte, error) {
	return json.Marshal(d)
}
