// Provides functions to handle creating and running an SSM document
package ssmrunner

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	asgInstanceTag = "aws:autoscaling:groupName"
)

type SSMCommand struct {
	BucketName     string
	deleteBucket   bool // Whether to delete the bucket on cleanup
	deleteDocument bool // whether to delete the document on cleanup
	Document       string
	Name           string
	Region         string
	Parameters     map[string]string
	init           bool // Whether the environment has been initialized
	Session        *session.Session
}

// Initialize the Environment for the SSM document to run
func (c *SSMCommand) Init() error {
	s3Svc := s3.New(c.Session)
	ssmSvc := ssm.New(c.Session)

	// Check for existing S3 Bucket
	headBucketInput := &s3.HeadBucketInput{
		Bucket: aws.String(c.BucketName),
	}
	_, headErr := s3Svc.HeadBucket(headBucketInput)
	if headErr == nil {
		// If S3 bucket exists assuming it should not be deleted
		c.deleteBucket = false
	} else {
		c.deleteBucket = true
		if headErr, ok := headErr.(awserr.Error); ok {
			switch headErr.Code() {
			// bucket not existing returns NotFound instead of NoSuchBucket
			case "NotFound":
				fallthrough
			case s3.ErrCodeNoSuchBucket:
				// Handle creation of bucket if it doesn't exist
				s3CreateBucketParams := &s3.CreateBucketInput{
					Bucket: aws.String(c.BucketName),
				}

				sessionRegion := aws.StringValue(c.Session.Config.Region)
				if sessionRegion != "us-east-1" {
					s3CreateBucketParams.CreateBucketConfiguration = &s3.CreateBucketConfiguration{
						LocationConstraint: aws.String(sessionRegion),
					}
				}
				_, err := s3Svc.CreateBucket(s3CreateBucketParams)
				if err != nil {
					return err
				}
			default:
				return headErr
			}
		} else {
			return headErr
		}
	}

	c.deleteDocument = false
	if c.Document != "" {
		createDocParams := &ssm.CreateDocumentInput{
			Content:      aws.String(c.Document),
			DocumentType: aws.String("Command"),
			Name:         aws.String(c.Name),
		}
		_, err := ssmSvc.CreateDocument(createDocParams)
		if err != nil {
			return err
		}

		c.deleteDocument = true
	}

	return nil
}

type RunInput struct {
	Instances      []string
	TagKey         string
	TagValue       string
	AutoScaleGroup string
}

func run(awsSess *session.Session, input *ssm.SendCommandInput) ([]ExecutedCommand, error) {
	ssmSvc := ssm.New(awsSess)

	sendResp, err := ssmSvc.SendCommand(input)
	if err != nil {
		return []ExecutedCommand{}, err
	}

	listInput := &ssm.ListCommandInvocationsInput{
		CommandId: sendResp.Command.CommandId,
	}
	// Sleep after invocation to ensure listCommand will return with something
	time.Sleep(2 * time.Second)

	cmdExecutions := make([]ExecutedCommand, 0)
	err = ssmSvc.ListCommandInvocationsPages(listInput,
		func(page *ssm.ListCommandInvocationsOutput, lastPage bool) bool {
			for _, invocation := range page.CommandInvocations {
				cmdExecutions = append(cmdExecutions, ExecutedCommand{
					CommandID:  aws.StringValue(sendResp.Command.CommandId),
					InstanceID: aws.StringValue(invocation.InstanceId),
				})
			}
			return true
		},
	)
	if err != nil {
		return []ExecutedCommand{}, err
	}

	return cmdExecutions, nil
}

func (c *SSMCommand) Run(ctx context.Context, input RunInput) ([]ExecutedCommand, error) {
	switch {
	case len(input.Instances) >= 1:
		return c.runInstances(input.Instances)
	case input.TagKey != "":
		if input.TagValue == "" {
			return []ExecutedCommand{}, fmt.Errorf("missing tag value")
		}
		return c.runTags(input.TagKey, input.TagValue)
	case input.AutoScaleGroup != "":
		return c.runAutoscaleGroup(input.AutoScaleGroup)
	default:
		return []ExecutedCommand{}, fmt.Errorf("no instances to run against")
	}
}

func (c *SSMCommand) runTags(key, value string) ([]ExecutedCommand, error) {
	input := &ssm.SendCommandInput{
		DocumentName:       aws.String(c.Name),
		OutputS3BucketName: aws.String(c.BucketName),
		Targets: []*ssm.Target{{
			Key:    aws.String("tag:" + key),
			Values: aws.StringSlice([]string{value}),
		}},
	}

	if len(c.Parameters) > 0 {
		params := make(map[string][]*string)
		for key, val := range c.Parameters {
			params[key] = aws.StringSlice([]string{val})
		}

		input.Parameters = params
	}

	return run(c.Session, input)
}

func (c *SSMCommand) runAutoscaleGroup(groupName string) ([]ExecutedCommand, error) {
	return c.runTags(asgInstanceTag, groupName)
}

// Run the command on the provided list of instances
func (c *SSMCommand) runInstances(instances []string) ([]ExecutedCommand, error) {
	input := &ssm.SendCommandInput{
		DocumentName:       aws.String(c.Name),
		InstanceIds:        aws.StringSlice(instances),
		OutputS3BucketName: aws.String(c.BucketName),
	}

	if len(c.Parameters) > 0 {
		params := make(map[string][]*string)
		for key, val := range c.Parameters {
			params[key] = aws.StringSlice([]string{val})
		}

		input.Parameters = params
	}

	return run(c.Session, input)
}

// Delete a bucket including all data in it
func deleteBucket(session *session.Session, bucketName string) error {
	svc := s3.New(session)

	lsObjectsInput := &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	}
	lsObjRes, err := svc.ListObjects(lsObjectsInput)
	if err != nil {
		return err
	}

	// Generate list of objects to delete from the bucket
	objDelList := make([]*s3.ObjectIdentifier, 0)

	for _, obj := range lsObjRes.Contents {
		objName := aws.StringValue(obj.Key)
		objDelList = append(objDelList, &s3.ObjectIdentifier{
			Key: aws.String(objName),
		})
	}

	delObjInput := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{
			Objects: objDelList,
			Quiet:   aws.Bool(true),
		},
	}
	_, err = svc.DeleteObjects(delObjInput)
	if err != nil {
		return err
	}

	delBucketInput := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	_, err = svc.DeleteBucket(delBucketInput)
	if err != nil {
		return err
	}

	return nil
}

// Cleanup environment for SSM Document
func (c *SSMCommand) Cleanup() error {
	ssmSvc := ssm.New(c.Session)

	if c.deleteDocument {
		deleteDocParams := &ssm.DeleteDocumentInput{
			Name: aws.String(c.Name),
		}

		_, err := ssmSvc.DeleteDocument(deleteDocParams)
		if err != nil {
			return err
		}
	}

	if c.deleteBucket {
		deleteBucket(c.Session, c.BucketName)
	}
	return nil
}
