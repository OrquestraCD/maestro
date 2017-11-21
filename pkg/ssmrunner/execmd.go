// Provides functions to handle creating and running an SSM document
package ssmrunner

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	OutputS3KeyPrefix = `[a-zA-Z0-9]+\/i-[a-zA-Z0-9]+\/[a-zA-Z0-9\.]+\/[a-zA-Z0-9\.]+\/`
	PollInterval      = 30
	S3UrlRegex        = `^https:\/\/s3(\.|\-)([a-z]+\-[a-z]+\-[0-9])?.amazonaws.com\/([a-zA-Z\-\_0-9]+)\/([a-zA-Z\-\_\/0-9\.\%]+)$`
)

type ExecutionError struct {
	OriginalError error
	InstanceID    string
	CommandID     string
}

func (e ExecutionError) Error() string {
	return fmt.Sprintf("(%s:%s) %s", e.InstanceID, e.CommandID, e.OriginalError)
}

type ExecutionErrors struct {
	Errors []ExecutionError
}

func NewExecutionErrors() ExecutionErrors {
	return ExecutionErrors{
		Errors: make([]ExecutionError, 0),
	}
}

func (e *ExecutionErrors) Add(err ExecutionError) {
	e.Errors = append(e.Errors, err)
}

func (e ExecutionErrors) Error() string {
	var err string

	for _, execErr := range e.Errors {
		err += execErr.Error() + "\n"
	}

	return err
}

func PollExecutedCommands(ctx context.Context, awsSess *session.Session, cmds []ExecutedCommand, output chan CommandOutput) error {
	execErrors := NewExecutionErrors()
	retErrors := make(chan ExecutionError, 5)

	go func() {
		for retError := range retErrors {
			execErrors.Add(retError)
		}
	}()

	var wg sync.WaitGroup
	for _, execCmd := range cmds {
		wg.Add(1)
		go func(cmd ExecutedCommand) {
			cmd.Poll(ctx, awsSess, output, retErrors)
			wg.Done()
		}(execCmd)
	}

	wg.Wait()
	close(retErrors)
	close(output)

	if len(execErrors.Errors) != 0 {
		return execErrors
	}

	return nil
}

type ExecutedCommand struct {
	CommandID   string
	CommandName string
	CommandOutput
	InstanceID string
}

// Cancel the running SSM Command on the instance.
func (e *ExecutedCommand) Cancel() {
	// TODO: Implement
}

// Check the finish statuses and return whether a bool
// indicating whether the command has finished.
func commandInvocationFinished(status string) bool {
	finishedStatuses := []string{
		"Delivery Timed Out",
		"Execution Timed Out",
		"Undeliverable",
		"Terminated",
		"Failed",
		"Success",
	}

	for _, finStatus := range finishedStatuses {
		if finStatus == status {
			return true
		}
	}

	return false
}

// Given the SSM Output bucket and original key returned by the API
// determine the actual SSM Output paths if they exist and return them.
func getOutputPaths(bucket, originalKey string, session *session.Session) ([]string, []string) {
	stdout := make([]string, 0)
	stderr := make([]string, 0)
	outputPrefix := strings.Join(strings.Split(originalKey, "/")[0:2], "/")

	s3Svc := s3.New(session)
	s3ListInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(outputPrefix),
	}
	// For now return empty if error is encountered
	s3Resp, err := s3Svc.ListObjectsV2(s3ListInput)
	if err != nil {
		return stdout, stderr
	}
	stdoutRe := regexp.MustCompile(OutputS3KeyPrefix + "stdout")
	stderrRe := regexp.MustCompile(OutputS3KeyPrefix + "stderr")

	var contentValue string
	for _, content := range s3Resp.Contents {
		contentValue = aws.StringValue(content.Key)
		if stdoutRe.MatchString(contentValue) {
			stdout = append(stdout, contentValue)
		}

		if stderrRe.MatchString(contentValue) {
			stderr = append(stderr, contentValue)
		}
	}

	return stdout, stderr
}

// Given an S3 URL return the bucket name and key
func getBucketInfo(url string) (string, string) {
	re := regexp.MustCompile(S3UrlRegex)
	result := re.FindStringSubmatch(url)

	return result[3], result[4]
}

// Poll the executedcommand to get the output
// TODO: Return the output responses in a channel rather than updating the Struct
func (e *ExecutedCommand) Poll(ctx context.Context, awsSess *session.Session, out chan CommandOutput, err chan ExecutionError) {
	var bucket, originalPath string
	commandReturned := false
	cmdOutput := CommandOutput{
		InstanceID: e.InstanceID,
	}

	ssmSvc := ssm.New(awsSess)
	for !commandReturned {
		select {
		// Recieved cancellation
		case <-ctx.Done():
			execErr := ExecutionError{
				OriginalError: ctx.Err(),
				InstanceID:    e.InstanceID,
				CommandID:     e.CommandID,
			}
			err <- execErr
			return
		default:
			cmdInvocInput := &ssm.GetCommandInvocationInput{
				CommandId:  aws.String(e.CommandID),
				InstanceId: aws.String(e.InstanceID),
			}

			resp, getErr := ssmSvc.GetCommandInvocation(cmdInvocInput)
			if getErr != nil {
				execErr := ExecutionError{
					OriginalError: getErr,
					InstanceID:    e.InstanceID,
					CommandID:     e.CommandID,
				}
				err <- execErr
				return
			}

			if commandInvocationFinished(aws.StringValue(resp.Status)) {
				commandReturned = true
				cmdOutput.Status = aws.StringValue(resp.Status)
				cmdOutput.CommandLog = aws.StringValue(resp.StandardOutputContent)

				stdoutResp := aws.StringValue(resp.StandardOutputUrl)
				if stdoutResp != "" {
					bucket, originalPath = getBucketInfo(stdoutResp)
				} else {
					stderrResp := aws.StringValue(resp.StandardErrorUrl)
					if stderrResp != "" {
						bucket, originalPath = getBucketInfo(stderrResp)
					}
				}

				if originalPath != "" {
					cmdOutput.Bucket = bucket
					cmdOutput.Stdout, cmdOutput.Stderr = getOutputPaths(bucket, originalPath, awsSess)
				}

				out <- cmdOutput
			} else {
				time.Sleep(PollInterval)
			}
		}
	}
}
