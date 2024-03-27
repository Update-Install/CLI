package commands

import (
	"ui/cli/module"

	"github.com/urfave/cli/v2"
)

func Config(*cli.Context) {
	module.GetConfigs()
}
