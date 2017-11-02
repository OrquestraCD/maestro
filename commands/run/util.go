package run

import (
	"strings"

	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/context"
	"github.com/rackerlabs/maestro/pkg/ec2utils"
	"github.com/rackerlabs/maestro/pkg/middleware"
	. "github.com/rackerlabs/maestro/ui"
)

func instanceList(c *cli.Context) []string {
	return strings.Split(c.String("instances"), ",")
}

// Detect Platform of a group of EC2 instances
func detectPlatform(c *cli.Context) (string, error) {
	UI.Debug("Attempting to determine the platform")

	context := context.New(middleware.GetSession(c))
	if c.String("instances") != "" {
		return ec2utils.DetectPlatformForInstances(
			context,
			instanceList(c),
		)
	}

	if c.String("autoscale-group") != "" {
		UI.Debug("Determining Autoscale Group Platform")
		return ec2utils.DetectPlatformForAutoscaleGroup(
			context,
			c.String("autoscale-group"),
		)
	}

	return ec2utils.DetectPlatformForTag(
		context,
		c.String("tag-key"),
		c.String("tag-value"),
	)
}
