package list

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rackerlabs/go-tables/tables"
	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/ec2utils"
	"github.com/rackerlabs/maestro/pkg/middleware"
	"github.com/rackerlabs/maestro/pkg/stringset"
	. "github.com/rackerlabs/maestro/ui"
)

const asgInstanceTag = "aws:autoscaling:groupName"

var defaultASGTblHeader = []string{
	"Name",
}

func listASGs(c *cli.Context) error {
	sess := middleware.GetSession(c)

	ec2Instances, err := ec2utils.DescribeInstances(sess, "")
	if err != nil {
		return err
	}
	UI.Debugf("DescribeInstances response %+v\n", ec2Instances)

	asgSet := stringset.New()
	for _, inst := range ec2Instances {
		for _, tag := range inst.EC2Information.Tags {
			if aws.StringValue(tag.Key) == asgInstanceTag {
				asgSet.Add(aws.StringValue(tag.Value))
			}
		}
	}

	tbl := tables.NewOrderedTable()
	tbl.AddColumn(append([]string{defaultASGTblHeader[0]}, asgSet.Slice()...))
	printOutput(tbl, strings.Split(c.String("fields"), ","))

	return nil
}
