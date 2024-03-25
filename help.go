package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func help(*cli.Context) error {
	fmt.Println("Update-Install tool for Linux")
	fmt.Println("Usage: update-install [command] [options]")
	fmt.Println("COMMANDS")
	fmt.Println("  help - show this help")
	fmt.Println("  install - install apps")
	return nil
}
