package ssmrunner

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	StdoutEmptyError = "Stdout is empty"
	StderrEmptyError = "Stderr is empty"
)

// Read a given S3 Document and return the body
func readS3Doc(bucket, key string, session *session.Session) (string, error) {
	svc := s3.New(session)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	retries := 5
	for i := 0; i < retries; i++ {
		resp, err := svc.GetObject(params)
		if err != nil {
			time.Sleep(30)
		} else {
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			return buf.String(), nil
		}
	}

	return "", fmt.Errorf("Unable to read key \"%s\" from bucket \"%s\"\n", key, bucket)
}

type CommandOutput struct {
	Bucket     string   // S3 Bucket the Output is stored in
	Stdout     []string // Full S3 Key for StdOut, empty if none exists
	Stderr     []string // Full S3 Key for Stderr, empty if none exists
	CommandLog string   // Any output provided by SSM
	InstanceID string   // InstanceID the output belongs to
	Status     string   // Status the SSM Run returned with
}

// Read a Command Stdout from S3 Bucket and return
// as a string.
// Pass AWS Session as input
func (c *CommandOutput) ReadStdout(session *session.Session) (string, error) {
	if len(c.Stdout) < 1 {
		return "", errors.New(StdoutEmptyError)
	}

	var out, document string
	var err error
	for _, doc := range c.Stdout {
		out, err = readS3Doc(c.Bucket, doc, session)
		if err != nil {
			return "", err
		}
		document += out
	}
	return document, nil
}

// Read a Command Stderr from S3 Bucket and return
// as a string.
// Pass AWS Session as input
func (c *CommandOutput) ReadStderr(session *session.Session) (string, error) {
	if len(c.Stderr) < 1 {
		return "", errors.New(StderrEmptyError)
	}

	var out, document string
	var err error
	for _, doc := range c.Stderr {
		out, err = readS3Doc(c.Bucket, doc, session)
		if err != nil {
			return "", err
		}
		document += out
	}
	return document, nil
}

// Given an error message return true if error
// message relates to Stdout being empty
func StdoutEmpty(err error) bool {
	if err != nil {
		return err.Error() == StdoutEmptyError
	}

	return false
}

// Given an error message return true if error
// message relates to Stderr being empty
func StderrEmpty(err error) bool {
	if err != nil {
		return err.Error() == StderrEmptyError
	}

	return false
}
