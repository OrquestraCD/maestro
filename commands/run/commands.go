package run

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"

	// Internal packages
	"github.com/rackerlabs/maestro/pkg/flagutil"
	"github.com/rackerlabs/maestro/pkg/stringid"
)

var documentFileFlag = cli.StringFlag{
	Name:  "file, f",
	Value: "",
	Usage: "The SSM Document file to upload and run.",
}

var documentNameFlag = cli.StringFlag{
	Name:  "name, n",
	Value: "",
	Usage: "The Name of the SSM Document to run.",
}

var parametersFlag = cli.StringFlag{
	Name:  "parameters, p",
	Value: "",
	Usage: "Parameters to pass to the SSM doc. (Key1=Value1 Key2=value2)",
}

var paramDelimiterFlag = cli.StringFlag{
	Name:  "parameters-delimiter, d",
	Value: " ",
	Usage: "Parameters delimiter to split on. (Ex with / Param=Value1/Param2=Value2)",
}

var instancesFlag = cli.StringFlag{
	Name:  "instances, i",
	Value: "",
	Usage: "Instance IDs to execute a given script/command on.",
}

var platformTypeFlag = cli.StringFlag{
	Name:  "platform, P",
	Value: "",
	Usage: "Specify what the platform type of the instances are. (ex: Linux, Windows)",
}

var noCleanupFlag = cli.BoolFlag{
	Name:  "no-clean, N",
	Usage: "Do not clean up the document and S3 bucket.",
}

var autoscaleFlag = cli.StringFlag{
	Name:  "autoscale-group, a",
	Value: "",
	Usage: "Autoscaling Group Name to execute command on.",
}

var tagKeyFlag = cli.StringFlag{
	Name:  "tag-key, K",
	Value: "",
	Usage: "Specify a tag key name for EC2 Instances. (Requires tag-value)",
}

var tagValueFlag = cli.StringFlag{
	Name:  "tag-value, V",
	Value: "",
	Usage: "Specify a tag value for EC2 Instances. (Requires tag-key)",
}

var bucketNameFlag = cli.StringFlag{
	Name:   "bucket-name, B",
	Value:  fmt.Sprintf("maestro-%s", stringid.GenerateID()),
	Usage:  "Name of the S3 Bucket to use for Maestro Output.",
	EnvVar: "MAESTRO_OUTPUT_BUCKET",
}

var Commands = []cli.Command{
	{
		Name:   "command",
		Usage:  "Runs a given command against specified instances.",
		Action: runShellCommand,
		Before: func(c *cli.Context) error {
			if len(c.Args()) == 0 {
				return errors.New("missing command argument")
			}

			groups := []flagutil.ExclusiveFlagGroup{
				[]string{"instances"},
				[]string{"autoscale-group"},
				[]string{"tag-key", "tag-value"},
			}

			return flagutil.ValidateExclusiveFlagGroups(c, groups)
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "alias, A",
				Usage: "Tells maestro the command being run is an alias set in the maestro config.",
			},
			autoscaleFlag,
			bucketNameFlag,
			instancesFlag,
			noCleanupFlag,
			platformTypeFlag,
			tagKeyFlag,
			tagValueFlag,
		},
	},
	{
		Name:   "document",
		Usage:  "Runs SSM Command document.",
		Action: runDocument,
		Before: func(c *cli.Context) error {
			if len(c.Args()) == 0 {
				return errors.New("missing document argument")
			}

			groups := []flagutil.ExclusiveFlagGroup{
				[]string{"instances"},
				[]string{"autoscale-group"},
				[]string{"tag-key", "tag-value"},
			}

			return flagutil.ValidateExclusiveFlagGroups(c, groups)
		},
		Flags: []cli.Flag{
			autoscaleFlag,
			bucketNameFlag,
			instancesFlag,
			noCleanupFlag,
			parametersFlag,
			paramDelimiterFlag,
			tagKeyFlag,
			tagValueFlag,
		},
	},
	{
		Name:  "script",
		Usage: "Runs a given script as an SSM document on provided instances",
		Before: func(c *cli.Context) error {
			if len(c.Args()) == 0 {
				return errors.New("missing script argument")
			}

			groups := []flagutil.ExclusiveFlagGroup{
				[]string{"instances"},
				[]string{"autoscale-group"},
				[]string{"tag-key", "tag-value"},
			}

			return flagutil.ValidateExclusiveFlagGroups(c, groups)
		},
		Action: runShellScript,
		Flags: []cli.Flag{
			autoscaleFlag,
			bucketNameFlag,
			instancesFlag,
			noCleanupFlag,
			platformTypeFlag,
			tagKeyFlag,
			tagValueFlag,
		},
	},
}
