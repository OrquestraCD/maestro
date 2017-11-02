package commands

import (
	"github.com/rackerlabs/maestro/commands/create"
	"github.com/rackerlabs/maestro/commands/list"
	"github.com/rackerlabs/maestro/commands/run"
	"github.com/rackerlabs/maestro/pkg/middleware"
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:        "create",
		Usage:       "Create SSM Documents locally",
		Subcommands: create.Commands,
	},
	{
		Name:        "list",
		Usage:       "Provides commands to list documents and instances",
		Subcommands: list.Commands,
		Before:      middleware.SetSession,
	},
	{
		Name:        "run",
		Usage:       "Provides subcommands for running SSM",
		Subcommands: run.Commands,
		Before:      middleware.SetSession,
	},
}
