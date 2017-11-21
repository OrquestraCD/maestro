package ec2utils

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Container struct for an instance
// Holds all the information returned by both SSM and EC2 APIs
type Instance struct {
	EC2Information ec2.Instance
	SSMInformation ssm.InstanceInformation
}

// Return SSM filters to be used in DescribeInstances if filters string
// is empty then only return static filter PingStatus.
func genSSMFilters(filters string) ([]*ssm.InstanceInformationStringFilter, error) {
	var filtersList []string
	if filters != "" {
		filtersList = strings.Split(filters, ",")
	}

	ssmFilters := make([]*ssm.InstanceInformationStringFilter, len(filtersList)+1)
	ssmFilters[0] = &ssm.InstanceInformationStringFilter{
		Key:    aws.String("PingStatus"),
		Values: []*string{aws.String("Online")},
	}

	if len(filtersList) == 0 {
		return ssmFilters, nil
	}

	i := 1
	for _, filter := range filtersList {
		filterInfo := strings.Split(filter, "=")

		if len(filterInfo) != 2 {
			return ssmFilters, fmt.Errorf("Invalid filter \"%s\", exect Key=Value", filter)
		}

		if filterInfo[0] == "PingStatus" {
			return ssmFilters, fmt.Errorf("PingStatus is static filter, cannot be altererd")
		}

		ssmFilters[i] = &ssm.InstanceInformationStringFilter{
			Key:    aws.String(filterInfo[0]),
			Values: []*string{aws.String(filterInfo[1])},
		}

		i++
	}

	return ssmFilters, nil
}

// Given SSM Filters return both the SSM description and the EC2 description
func DescribeInstances(session *session.Session, ssmFilters string) ([]Instance, error) {
	var err error
	results := make([]Instance, 0)
	instanceMap := make(map[string]Instance)
	ssmSvc := ssm.New(session)

	ssmInput := ssm.DescribeInstanceInformationInput{}
	ssmInput.Filters, err = genSSMFilters(ssmFilters)
	if err != nil {
		return results, err
	}

	instanceIDs := make([]*string, 0)
	err = ssmSvc.DescribeInstanceInformationPages(&ssmInput,
		func(page *ssm.DescribeInstanceInformationOutput, lastPage bool) bool {
			for _, inst := range page.InstanceInformationList {
				instanceIDs = append(instanceIDs, inst.InstanceId)
				instanceMap[aws.StringValue(inst.InstanceId)] = Instance{SSMInformation: *inst}
			}

			return true
		},
	)
	if err != nil {
		return results, err
	}

	// If no SSM instances are available return an empty slice
	if len(instanceIDs) == 0 {
		return []Instance{}, nil
	}

	ec2Svc := ec2.New(session)
	describeInstInput := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	}

	err = ec2Svc.DescribeInstancesPages(describeInstInput,
		func(page *ec2.DescribeInstancesOutput, latPage bool) bool {
			for _, reservation := range page.Reservations {
				for _, inst := range reservation.Instances {
					tempInst := instanceMap[aws.StringValue(inst.InstanceId)]
					tempInst.EC2Information = *inst
					instanceMap[aws.StringValue(inst.InstanceId)] = tempInst
				}
			}
			return true
		},
	)
	if err != nil {
		return results, err
	}

	for _, inst := range instanceMap {
		results = append(results, inst)
	}

	return results, nil
}
