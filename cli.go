package main

import (
	"fmt"
	"internal/commands"

	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
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
			},
			Action: func(c *cli.Context) error {

				return commands.HandleBuildCommand(c.Args().First(), c.Bool("dts"))
			},
		},
		{
			Name:  "dts",
			Usage: "Emit .d.ts files for package",
			Action: func(c *cli.Context) error {
				return commands.RunDts()
			},
		},
		{
			Name:  "prettier",
			Usage: "Will run pretty-quick",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "all",
					Usage: "Will prettify all files instead of staged files",
				},
			},
			Action: func(c *cli.Context) error {
				return commands.HandlePrettierCommand(c.Bool("all"))
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
		{
			Name:    "version",
			Usage:   "Version of cli",
			Aliases: []string{"v"},
			Action: func(c *cli.Context) error {
				fmt.Printf("tsdev %s, commit %s, built at %s by %s", version, commit, date, builtBy)
				return nil
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

			if c.Bool("version") {
				fmt.Printf("tsdev %s, commit %s, built at %s by %s", version, commit, date, builtBy)
				return nil
			} else {
				return commands.HandleDefault(c.Bool("watch"), c.Args().Slice())
			}
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "watch",
				Usage: "Run in watch mode",
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
			},
		},
	}

	return *app, nil
}
