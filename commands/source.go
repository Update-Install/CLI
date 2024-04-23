package commands

import (
	"log"
	"ui/cli/module"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

func Source(c *cli.Context) {
	if c.Args().Get(0) == "list" || (c.String("name") == "") || (c.String("url") == "") {
		config, err := module.GetSource()
		if err != nil {
			log.Fatal(err)
		}

		if len(config.Source) == 0 {
			color.Yellowln("No files are installed.")
			return
		}

		for _, file := range config.Source {
			color.Cyanf("%s: ", file.Name)
			color.Greenln(file.URL)
		}

		return
	}

	config, err := module.SetSource(c.String("name"), c.String("url"))
	if err != nil {
		log.Fatal(err)
	}

	module.SaveToConfigFile(config)
}
