package commands

import (
	"log"
	"ui/cli/module"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

func Config(c *cli.Context) {

	if (c.String("name") == "") || (c.String("url") == "") {
		config, err := module.GetConfigs()
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range config.Files {
			color.Cyanf("%s: ", file.Name)
			color.Greenln(file.URL)
		}

		return
	}

	config, err := module.SetFileURLConfigs(c.String("name"), c.String("url"))
	if err != nil {
		log.Fatal(err)
	}

	module.SaveToConfigFile(config)
}
