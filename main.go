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
		Name:    "UI-CLI",
		Usage:   "A app install/update tool for Linux",
		Version: "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Version of CLI",
				Action: func(ctx *cli.Context) error {
					fmt.Println(ctx.App.Version)
					return nil
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install package",
				Action: func(ctx *cli.Context) error {
					commands.Install(ctx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
