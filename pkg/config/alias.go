package config

import (
	"encoding/json"
	"fmt"

	"github.com/rackerlabs/maestro/pkg/scripts"
	"github.com/rackerlabs/maestro/pkg/ssmdoc"
)

const missingAttrText = "%s must be set for alias %s"

type Alias struct {
	Description string `json:"description"`
	Command     string `json:"command,omitempty"`
	Name        string
	Platform    string `json:"platform,omitempty"`
	Type        string `json:"type"`
}

func stringMapValues(hash map[string]string) []string {
	values := make([]string, len(hash))
	i := 0
	for _, val := range hash {
		values[i] = val
		i++
	}

	return values
}

// Validate Alias is correct
func (a Alias) validate() error {
	if a.Command == "" {
		return fmt.Errorf(missingAttrText, "command", a.Name)
	}

	return nil
}

func (a Alias) Document() (*[]byte, error) {
	doc, err := ssmdoc.NewDocument(a.Description)
	if err != nil {
		return nil, err
	}

	ssmRunCommand := scripts.ScriptToSSMCommands(a.Command)
	scriptType := scripts.ScriptTypeByName(a.Type)

	switch scriptType {
	case scripts.PowerShell:
		runPowerShellInput := ssmdoc.RunPowerShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(ssmRunCommand),
		}

		runPowerShell := ssmdoc.Plugin{
			Action: ssmdoc.RunPowerShellScriptPluginAction,
			Name:   "runPowerShellScript",
			Inputs: runPowerShellInput,
		}

		if err := doc.AddStep(runPowerShell); err != nil {
			return nil, err
		}
	case scripts.Bash:
		runShellProps := ssmdoc.RunShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(ssmRunCommand),
		}

		runShell := ssmdoc.Plugin{
			Action: ssmdoc.RunShellScriptPluginAction,
			Name:   "runShellScript",
			Inputs: runShellProps,
		}

		if err := doc.AddStep(runShell); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unable to determine script %s type.\n", a.Name)
	}

	encodedDoc, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	return &encodedDoc, nil
}
