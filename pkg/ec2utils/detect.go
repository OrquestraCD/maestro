package ec2utils

import (
	"fmt"

	"github.com/rackerlabs/maestro/pkg/context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Detect Platform Type of a group of EC2 instances
func DetectPlatformForInstances(ctx *context.Context, instances []string) (string, error) {
	input := ssm.DescribeInstanceInformationInput{
		Filters: []*ssm.InstanceInformationStringFilter{{
			Key:    aws.String("InstanceIds"),
			Values: aws.StringSlice(instances),
		}},
	}

	resp, err := ctx.SSMClient.DescribeInstanceInformation(&input)
	if err != nil {
		return "", err
	}

	if len(resp.InstanceInformationList) == 0 {
		return "", fmt.Errorf("none of the provided instances are configured with SSM")
	}

	inst := resp.InstanceInformationList[0]
	return aws.StringValue(inst.PlatformType), nil
}

// Return a Platform Type given a tag key name and value
func DetectPlatformForTag(ctx *context.Context, key, value string) (string, error) {
	ec2DescInput := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{{
			Name:   aws.String("tag:" + key),
			Values: aws.StringSlice([]string{value}),
		}},
	}

	ec2Resp, err := ctx.EC2Client.DescribeInstances(ec2DescInput)
	if err != nil {
		return "", err
	}

	if len(ec2Resp.Reservations) == 0 {
		return "", fmt.Errorf("no instances with key %s and value %s available", key, value)
	}

	instanceIds := make([]string, 0)
	for _, res := range ec2Resp.Reservations {
		for _, inst := range res.Instances {
			instanceIds = append(instanceIds, aws.StringValue(inst.InstanceId))
		}
	}

	ssmInput := &ssm.DescribeInstanceInformationInput{
		Filters: []*ssm.InstanceInformationStringFilter{{
			Key:    aws.String("InstanceIds"),
			Values: aws.StringSlice(instanceIds),
		}},
	}

	ssmResp, err := ctx.SSMClient.DescribeInstanceInformation(ssmInput)
	if err != nil {
		return "", err
	}

	if len(ssmResp.InstanceInformationList) == 0 {
		return "", fmt.Errorf("no instances with key %s are configured with SSM", key)
	}

	platform := aws.StringValue(ssmResp.InstanceInformationList[0].PlatformType)
	return platform, nil
}

func DetectPlatformForAutoscaleGroup(ctx *context.Context, asg string) (string, error) {
	return DetectPlatformForTag(ctx, "aws:autoscaling:groupName", asg)
}
