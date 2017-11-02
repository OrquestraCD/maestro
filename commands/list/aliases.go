package list

import (
	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/middleware"
	. "github.com/rackerlabs/maestro/ui"
)

func listAliases(c *cli.Context) error {
	conf := middleware.GetMaestroConfig(c)
	for name, aObj := range conf.Aliases {
		description := "NO DESCRIPTION AVAILABLE"
		if aObj.Description != "" {
			description = aObj.Description
		}

		UI.Printf("%s - %s\n", name, description)
	}

	return nil
}
