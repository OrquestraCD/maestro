package ec2utils

import (
	"testing"

	"github.com/rackerlabs/maestro/pkg/context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockedEC2DescribeInstances struct {
	ec2iface.EC2API
	Resp *ec2.DescribeInstancesOutput
}

func (m mockedEC2DescribeInstances) DescribeInstances(in *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return m.Resp, nil
}

type mockedSSMDescribeInstanceInformation struct {
	ssmiface.SSMAPI
	Resp *ssm.DescribeInstanceInformationOutput
}

func (m mockedSSMDescribeInstanceInformation) DescribeInstanceInformation(in *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
	return m.Resp, nil
}

func TestDetectPlatformForInstances(t *testing.T) {
	instanceId := "i-k234lkj2l34"
	platformType := "Linux"
	ssmResp := &ssm.DescribeInstanceInformationOutput{
		InstanceInformationList: []*ssm.InstanceInformation{{
			InstanceId:   aws.String(instanceId),
			PlatformType: aws.String(platformType),
		}},
	}

	ctx := context.Context{
		EC2Client: mockedEC2DescribeInstances{Resp: nil},
		SSMClient: mockedSSMDescribeInstanceInformation{Resp: ssmResp},
	}

	got, err := DetectPlatformForInstances(&ctx, []string{"i-k234lkj2l34"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != platformType {
		t.Errorf("Expected \"%s\", Got \"%s\"", platformType, got)
	}
}

func TestDetectPlatformForTag(t *testing.T) {
	instanceId := "i-k234lkj2l34"
	platformType := "Linux"
	ssmResp := &ssm.DescribeInstanceInformationOutput{
		InstanceInformationList: []*ssm.InstanceInformation{{
			InstanceId:   aws.String(instanceId),
			PlatformType: aws.String(platformType),
		}},
	}

	ec2Resp := &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{{
			Instances: []*ec2.Instance{{
				InstanceId: aws.String(instanceId),
			}},
		}},
	}

	ctx := context.Context{
		EC2Client: mockedEC2DescribeInstances{Resp: ec2Resp},
		SSMClient: mockedSSMDescribeInstanceInformation{Resp: ssmResp},
	}

	got, err := DetectPlatformForTag(&ctx, "test-tag", "test-value")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != platformType {
		t.Errorf("Expected \"%s\", Got \"%s\"", platformType, got)
	}
}

func TestDetectPlatformForAutoscaleGroup(t *testing.T) {
	instanceId := "i-k234lkj2l34"
	platformType := "Linux"
	ssmResp := &ssm.DescribeInstanceInformationOutput{
		InstanceInformationList: []*ssm.InstanceInformation{{
			InstanceId:   aws.String(instanceId),
			PlatformType: aws.String(platformType),
		}},
	}

	ec2Resp := &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{{
			Instances: []*ec2.Instance{{
				InstanceId: aws.String(instanceId),
			}},
		}},
	}

	ctx := context.Context{
		EC2Client: mockedEC2DescribeInstances{Resp: ec2Resp},
		SSMClient: mockedSSMDescribeInstanceInformation{Resp: ssmResp},
	}

	got, err := DetectPlatformForAutoscaleGroup(&ctx, "myASG")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != platformType {
		t.Errorf("Expected \"%s\", Got \"%s\"", platformType, got)
	}
}
