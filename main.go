package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"ui/cli/commands"
)

func main() {
	app := &cli.App{
		Name:     "UI-CLI",
		HelpName: "ui",
		Usage:    "Simple tool for installing deb package for LInux Distributions.",
		Version:  "v0.7.3",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version of ui",
				Action: func(ctx *cli.Context) error {
					fmt.Println(ctx.App.Version)
					return nil
				},
			},
			{
				Name:        "install",
				Aliases:     []string{"i"},
				Usage:       "install or update packages",
				Description: "If no package name is specified, it will install all packages in configured source.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "name of the file to install",
					},
				},
				Action: func(ctx *cli.Context) error {
					commands.Install(ctx)
					return nil
				},
			},
			{
				Name:        "source",
				Aliases:     []string{"s"},
				Usage:       "manage source of package",
				Description: "If no options are specified, it will list all sources of package on local machine.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "package name of the source",
					},
					&cli.StringFlag{
						Name:    "url",
						Aliases: []string{"u"},
						Usage:   "url of the source",
					},
				},
				Action: func(ctx *cli.Context) error {
					commands.Source(ctx)
					return nil
				},
			},
			{
				Name:    "update-self",
				Aliases: []string{"us"},
				Usage:   "update ui-cli to the latest version",
				Action: func(ctx *cli.Context) error {
					commands.UpdateSelf(ctx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
