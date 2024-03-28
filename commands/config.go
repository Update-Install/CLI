package commands

import (
	"fmt"
	"log"
	"ui/cli/module"

	"github.com/urfave/cli/v2"
)

func Config(c *cli.Context) {

	if (c.String("name") == "") || (c.String("url") == "") {
		config, err := module.GetConfigs()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(config)
		return
	}

	config, err := module.SetFileURLConfigs(c.String("name"), c.String("url"))
	if err != nil {
		log.Fatal(err)
	}

	module.SaveToConfigFile(config)
}
