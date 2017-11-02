package list

import "github.com/urfave/cli"

var Commands = []cli.Command{
	{
		Name:   "aliases",
		Usage:  "List all aliases in the Maestro configuration file.",
		Action: listAliases,
		Before: func(c *cli.Context) error {
			return nil
		},
	},
	{
		Name:   "asgs",
		Usage:  "List asgs available to AWS SSM",
		Action: listASGs,
		Before: func(c *cli.Context) error {
			return validateFieldNames(c, defaultASGTblHeader)
		},
		Flags: []cli.Flag{
			fieldsFlag(defaultASGTblHeader),
		},
	},
	{
		Name:   "documents",
		Usage:  "List documents available in AWS SSM",
		Action: listSSMDocuments,
		Before: func(c *cli.Context) error {
			return validateFieldNames(c, defaultDocTblHeader)
		},
		Flags: []cli.Flag{
			fieldsFlag(defaultDocTblHeader),
		},
	},
	{
		Name:   "instances",
		Usage:  "List instances available for SSM.",
		Action: listInstancesCli,
		Before: func(c *cli.Context) error {
			return validateFieldNames(c, defaultInstTblHeader)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "filters,F",
				Usage: "Filter instances based on key values. (eg: PlatformTypes=Linux)",
				Value: "",
			},
			fieldsFlag(defaultInstTblHeader),
			cli.BoolFlag{
				Name:  "list,l",
				Usage: "Print a comma delimited list of instances.",
			},
		},
	},
}
