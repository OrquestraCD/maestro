package run

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/middleware"
	"github.com/rackerlabs/maestro/pkg/ssmrunner"
	. "github.com/rackerlabs/maestro/ui"
)

func loadParameters(params, delimiter string) (map[string]string, error) {
	paramList := strings.Split(params, delimiter)
	result := make(map[string]string)

	for _, param := range paramList {
		args := strings.Split(param, "=")
		if len(args) != 2 {
			return result, fmt.Errorf("Not enough arguments for paramter.\n")
		}

		result[args[0]] = args[1]
	}

	return result, nil
}

func runDocument(c *cli.Context) error {
	sess := middleware.GetSession(c)

	parameters := make(map[string]string)
	var err error
	if c.String("parameters") != "" {
		parameters, err = loadParameters(
			c.String("parameters"),
			c.String("parameters-delimiter"),
		)

		if err != nil {
			return err
		}
		UI.Debugf("Parsed Parameters \"%+v\"\n", parameters)
	}

	command := ssmrunner.SSMCommand{
		BucketName: c.String("bucket-name"),
		Name:       c.Args()[0],
		Parameters: parameters,
		Session:    sess,
	}
	// Create the Command and initialize the environment
	if err := command.Init(); err != nil {
		return err
	}

	if !c.Bool("no-clean") {
		defer func() {
			if err := command.Cleanup(); err != nil {
				UI.Debugf("Error occurred while cleaning up: %v\n", err)
			}
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

	for out := range output {
		printOutput(out, sess)
	}

	return nil
}
