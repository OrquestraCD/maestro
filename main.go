package main

import (
	"fmt"
	"os"

	"github.com/rackerlabs/maestro/commands"
	"github.com/rackerlabs/maestro/pkg/config"
	"github.com/rackerlabs/maestro/pkg/middleware"
	"github.com/rackerlabs/maestro/ui"

	"github.com/urfave/cli"
)

var Version = ""

func main() {
	defaultConfig := config.DefaultConfig()

	app := cli.NewApp()
	app.Name = "maestro"
	app.EnableBashCompletion = true
	app.Usage = "Improved SSM CLI."
	app.Metadata = make(map[string]interface{})
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, D",
			Usage:  "Turn on Debug mode for verbose output",
			EnvVar: "MAESTRO_DEBUG",
		},
		cli.StringFlag{
			Name:   "region, r",
			Usage:  "AWS Region (optional)",
			Value:  "",
			EnvVar: "MAESTRO_REGION",
		},
		cli.StringFlag{
			Name:   "config, c",
			Usage:  fmt.Sprintf("Path to a maestro config, defaults to %s", defaultConfig),
			EnvVar: "MAESTRO_CONFIG",
			Value:  defaultConfig,
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			ui.InitUI(ui.Debug)
			ui.UI.Debug("Debug mode set")

			// Print the values of the global flags
			for _, flg := range []string{"region"} {
				ui.UI.Debugf("%s set to %+v\n", flg, c.GlobalString(flg))
			}
		} else {
			ui.InitUI(ui.Standard)
		}

		// TODO: Load Config
		ui.UI.Debugf("Attempting to load Maestro Config %s\n", c.GlobalString("config"))
		if _, err := os.Stat(c.GlobalString("config")); os.IsNotExist(err) {
			ui.UI.Debug("No configuration file loaded ... using empty default.")
		} else {
			if err := middleware.SetMaestroConfig(c); err != nil {
				return err
			}
			ui.UI.Debug("Config loaded")
		}

		middleware.SetAWSConfig(c)
		return nil
	}

	app.Commands = commands.Commands
	if err := app.Run(os.Args); err != nil {
		ui.UI.Error(err.Error())
		os.Exit(1)
	}
}
