// Commands to generate and run new temporary SSM documents
package run

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/fileutil"
	"github.com/rackerlabs/maestro/pkg/middleware"
	"github.com/rackerlabs/maestro/pkg/scripts"
	"github.com/rackerlabs/maestro/pkg/ssmdoc"
	"github.com/rackerlabs/maestro/pkg/ssmrunner"
	. "github.com/rackerlabs/maestro/ui"
)

// Generate a new SSM Document given a specific command and run
func createAndRunCommand(cmd []string, c *cli.Context) error {

	// Generate new SSM document with command
	doc, err := ssmdoc.NewDocument("Temporary Command Document")
	if err != nil {
		return err
	}

	platform := c.String("platform")
	if platform == "" {
		// Attempt to detect the platform if one is not set
		platform, err = detectPlatform(c)
		if err != nil {
			return err
		}
	}

	switch platform {
	case "Windows":
		UI.Debug("Platform Windows selected")
		runPowerShellInput := ssmdoc.RunPowerShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(cmd),
		}
		runPowerShell := ssmdoc.Plugin{
			Action: ssmdoc.RunPowerShellScriptPluginAction,
			Name:   "runPowerShellScript",
			Inputs: runPowerShellInput,
		}

		err := doc.AddStep(runPowerShell)
		if err != nil {
			return err
		}
	case "Linux":
		UI.Debug("Platform Linux selected")
		runShellInput := ssmdoc.RunShellScriptPluginInput{
			RunCommand: ssmdoc.ListValue(cmd),
		}
		runShell := ssmdoc.Plugin{
			Action: ssmdoc.RunShellScriptPluginAction,
			Name:   "runShellScript",
			Inputs: runShellInput,
		}

		err := doc.AddStep(runShell)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid platform %s\n", platform)
	}

	document, err := doc.String()
	if err != nil {
		return err
	}
	UI.Debugf("Document created: %s\n", document)

	sess := middleware.GetSession(c)
	docName := c.String("name")
	now := time.Now().Unix()
	if docName == "" {
		docName = fmt.Sprintf("%d-temp-document", now)
	}

	// Create the Command and initialize the environment
	command := ssmrunner.SSMCommand{
		BucketName: c.String("bucket-name"),
		Document:   document,
		Parameters: map[string]string{},
		Session:    sess,
		Name:       docName,
	}

	if err := command.Init(); err != nil {
		return err
	}

	if !c.Bool("no-clean") {
		defer func() {
			UI.Debug("Cleaning up environment")
			if err := command.Cleanup(); err != nil {
				UI.Debugf("Error occurred while cleaning up: %v\n", err)
			}
			UI.Debug("Clean up complete")
		}()
	}

	runInput := ssmrunner.RunInput{}
	switch {
	case c.String("autoscale-group") != "":
		runInput.AutoScaleGroup = c.String("autoscale-group")
	case c.String("tag-key") != "":
		runInput.TagKey = c.String("tag-key")
		runInput.TagValue = c.String("tag-value")
	default:
		instanceIDs := strings.Split(c.String("instances"), ",")
		runInput.Instances = instanceIDs
	}

	executions, err := command.Run(context.Background(), runInput)
	if err != nil {
		return err
	}

	output := make(chan ssmrunner.CommandOutput, len(executions)+1)
	if err := ssmrunner.PollExecutedCommands(context.Background(), sess, executions, output); err != nil {
		return err
	}

	// Print the output
	for out := range output {
		printOutput(out, sess)
	}

	return nil
}

// Execute a Command against n instances with SSM
func runShellCommand(c *cli.Context) error {
	var cmd []string
	if c.Bool("alias") {
		UI.Debugf("Alias flag set, looking up alias")

		conf := middleware.GetMaestroConfig(c)
		aliasName := c.Args()[0]
		aliasCmd, ok := conf.Aliases[aliasName]
		if !ok {
			return fmt.Errorf("alias %s not found\n", aliasName)
		}

		cmd = []string{aliasCmd.Command}
		if aliasCmd.Platform != "" {
			c.Set("platform", aliasCmd.Platform)
		}
	} else {
		cmd = []string{strings.Join(c.Args(), " ")}
	}

	UI.Debugf("Running command \"%s\"\n", cmd)

	return createAndRunCommand(cmd, c)
}

// Execute a script against n instances with SSM
func runShellScript(c *cli.Context) error {
	scr, err := fileutil.ReadFileToString(c.Args()[0])
	if err != nil {
		return err
	}
	ssmRunCommands := scripts.ScriptToSSMCommands(scr)

	return createAndRunCommand(ssmRunCommands, c)
}
