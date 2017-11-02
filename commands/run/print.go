package run

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/rackerlabs/maestro/pkg/ssmrunner"
	. "github.com/rackerlabs/maestro/ui"
)

// Print stdout and stderr in a specif format
func printInstanceOutput(instanceID, status, stdout, stderr string) {
	header := fmt.Sprintf("######################%s######################", instanceID)
	UI.Print(header)

	if stdout != "" {
		UI.Print(stdout)
	}

	if stderr != "" {
		UI.Print(stderr)
	}

	UI.Printf("Exited with status: %s\n", status)
}

// Print Output from running a command
func printOutput(output ssmrunner.CommandOutput, session *session.Session) error {
	stdout, err := output.ReadStdout(session)
	if err != nil {
		if ssmrunner.StdoutEmpty(err) {
			if len(output.CommandLog) == 0 {
				stdout = "Stdout empty"
			} else {
				stdout = output.CommandLog
			}
		} else {
			UI.Debugf("Received error reading from stdout: %v\n", err)
		}
	}

	stderr, err := output.ReadStderr(session)
	if err != nil && !ssmrunner.StderrEmpty(err) {
		UI.Debugf("Recieved error from stderr reader: %v\n", err)
	}

	printInstanceOutput(output.InstanceID, output.Status, stdout, stderr)
	return nil
}
