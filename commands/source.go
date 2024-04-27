package commands

import (
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Source(c *cli.Context) error {
	if c.Args().Get(0) == "list" || (c.String("name") == "") || (c.String("url") == "") {
		config, err := module.GetSource()
		if err != nil {
			return err
		}

		if len(config.Source) == 0 {
			color.Yellowln("There are no package in source.")
			return nil
		}

		for _, file := range config.Source {
			color.Cyanf("%s: ", file.Name)
			color.Greenln(file.URL)
		}

		return nil
	}

	config, err := module.SetSource(c.String("name"), c.String("url"))
	if err != nil {
		return err
	}

	module.SaveToConfigFile(config)
	return nil
}
