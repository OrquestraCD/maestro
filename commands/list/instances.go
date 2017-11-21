package list

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rackerlabs/go-tables/tables"
	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/ec2utils"
	"github.com/rackerlabs/maestro/pkg/middleware"
	. "github.com/rackerlabs/maestro/ui"
)

var defaultInstTblHeader = []string{
	"Instance ID",
	"Name",
	"Agent Version",
	"Platform Type",
	"Ping Status",
	"Latest",
}

func listInstancesCli(c *cli.Context) error {
	sess := middleware.GetSession(c)

	ec2Instances := make([][]string, 1)
	ec2Instances[0] = defaultInstTblHeader
	instances, err := ec2utils.DescribeInstances(sess, c.String("filters"))
	if err != nil {
		return err
	}

	UI.Debugf("DescribeInstances response: %+v\n", instances)
	// If List flag is set then print a comma delimited list of instances
	if c.Bool("list") {
		instanceIDs := make([]string, len(instances))
		for i, inst := range instances {
			instanceIDs[i] = aws.StringValue(inst.SSMInformation.InstanceId)
		}
		UI.Print(strings.Join(instanceIDs, ","))
		return nil
	}

	var name string
	for _, inst := range instances {
		// Add the id to a soon to be used filter for describe ec2 instances
		name = ""

		// Find the Name tag and assign it to name
		for _, tag := range inst.EC2Information.Tags {
			if aws.StringValue(tag.Key) == "Name" {
				name = aws.StringValue(tag.Value)
				break
			}
		}

		ec2Instances = append(ec2Instances, []string{
			aws.StringValue(inst.SSMInformation.InstanceId),
			name, //Placeholder for name
			aws.StringValue(inst.SSMInformation.AgentVersion),
			aws.StringValue(inst.SSMInformation.PlatformType),
			aws.StringValue(inst.SSMInformation.PingStatus),
			fmt.Sprintf("%v", aws.BoolValue(inst.SSMInformation.IsLatestVersion)),
		})
	}

	tbl := tables.NewOrderedTableFromMatrix(ec2Instances)
	printOutput(tbl, strings.Split(c.String("fields"), ","))

	return nil
}
