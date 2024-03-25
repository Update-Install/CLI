package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "UI-CLI",
		Usage: "A app install/update tool for Linux",
		Commands: []*cli.Command{
			{
				Name:    "help",
				Aliases: []string{"i"},
				Usage:   "help",
				Action: func(c *cli.Context) error {
					return help(c)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
