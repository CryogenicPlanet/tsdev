package main

import (
	"internal/commands"

	"github.com/urfave/cli/v2"
)

func SetupCliApp() (cli.App, error) {

	cliCommands := []*cli.Command{
		{Name: "create",
			Usage: "Create a new application",
			Action: func(c *cli.Context) error {
				return commands.HandleCreateCommand(c.Args().First())
			},
		},
		{
			Name:  "dev",
			Usage: "This run the app in dev mode with file watching",
			Action: func(c *cli.Context) error {
				//  handle dev
				return commands.HandleDevCommand(c.Args().First())
			},
		},
		{
			Name:  "build",
			Usage: "This builds the app for production.",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "dts",
					Usage: "Will emit .d.ts files and bundle them",
				},
				// &cli.BoolFlag{
				// 	Name:  "sourcemap",
				// 	Usage: "Will emit source maps",
				// },
				// &cli.StringFlag{
				// 	Name:  "dist",
				// 	Usage: "Set output directory",
				// 	Value: "dist",
				// },
			},
			Action: func(c *cli.Context) error {

				return commands.HandleBuildCommand(c.Args().First(), c.Bool("dts"))
			},
		},
		{
			Name:  "prettier",
			Usage: "Will run pretty-quick",
			Action: func(c *cli.Context) error {
				return commands.HandlePrettierCommand()
			},
		},
		{
			Name:  "lint",
			Usage: "Will lint the application",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "fix",
					Usage: "Will auto fix linter problems",
				},
			},
			Action: func(c *cli.Context) error {
				return commands.HandleLintCommand(c.Bool("fix"))
			},
		},
	}

	app := &cli.App{
		Name:                 "tsdev",
		Commands:             cliCommands,
		Usage:                "Zero config modern typescript tooling",
		EnableBashCompletion: true,
		ArgsUsage:            "Run a .ts file with zero config directly",
		Action: func(c *cli.Context) error {

			return commands.HandleDefault(c.Bool("watch"), c.Args().Slice())
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "watch",
				Usage: "Run in watch mode",
			},
		},
	}

	return *app, nil
}
