package create

import (
	"errors"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:   "command",
		Usage:  "Create a runShell or runPowershell SSM document from a provided script.",
		Action: generateRunCommand,
		Before: func(c *cli.Context) error {
			if len(c.Args()) == 0 {
				return errors.New("Missing script argument")
			}

			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "Name of the resulting SSM file. Optional.",
			},
			cli.StringFlag{
				Name:  "type, t",
				Value: "",
				Usage: "Type of SSM Run Document to create. (ie: powershell, bash) Optional",
			},
		},
	},
}
