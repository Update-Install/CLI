package commands

import (
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(*cli.Context) {
	module.CreateCacheDotFolder()
}
