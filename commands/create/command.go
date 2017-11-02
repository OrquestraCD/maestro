package create

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/urfave/cli"

	// Internal Packages
	"github.com/rackerlabs/maestro/pkg/fileutil"
	"github.com/rackerlabs/maestro/pkg/json"
	"github.com/rackerlabs/maestro/pkg/scripts"
	"github.com/rackerlabs/maestro/pkg/ssmdoc"
	. "github.com/rackerlabs/maestro/ui"
)

func generateRunCommand(c *cli.Context) error {
	scriptName := path.Base(c.Args()[0])
	UI.Debugf("Attempting to opening file %s\n", scriptName)
	outputName := c.String("output")
	if outputName == "" {
		outputName = scriptName + ".json"
	}

	documentDescription := fmt.Sprintf(
		"Generated from script %s on %s",
		scriptName,
		time.Now(),
	)

	scriptContent, err := fileutil.ReadFileToString(scriptName)
	if err != nil {
		return err
	}
	ssmRunCommand := scripts.ScriptToSSMCommands(scriptContent)

	doc, err := ssmdoc.NewDocument(documentDescription)
	if err != nil {
		return err
	}

	scriptType := scripts.DetectScriptByExtension(scriptName)
	switch scriptType {
	case scripts.PowerShell:
		UI.Debug("Detected powershell script")
		runPowerShellInput := ssmdoc.RunPowerShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(ssmRunCommand),
		}
		runPowerShell := ssmdoc.Plugin{
			Action: ssmdoc.RunPowerShellScriptPluginAction,
			Name:   "runPowerShellScript",
			Inputs: runPowerShellInput,
		}

		if err := doc.AddStep(runPowerShell); err != nil {
			return err
		}
	case scripts.Bash:
		UI.Debug("Detected bash script")
		runShellProps := ssmdoc.RunShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(ssmRunCommand),
		}
		runShell := ssmdoc.Plugin{
			Action: ssmdoc.RunShellScriptPluginAction,
			Name:   "runShellScript",
			Inputs: runShellProps,
		}

		if err := doc.AddStep(runShell); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unable to detect script type set \"--type\" flag")
	}

	document, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	UI.Debugf("Document generated: %s\n", string(document))

	if err := ioutil.WriteFile(outputName, document, 0750); err != nil {
		return err
	}

	return nil
}
