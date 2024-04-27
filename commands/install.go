package commands

import (
	"errors"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(c *cli.Context) error {
	var cmd = exec.Command("sudo", "-v")
	cmd.Stdout = nil
	err := cmd.Run()
	if err != nil {
		return err
	}

	module.CreateCacheDotFolder()
	config, err := module.GetSource()
	if err != nil {
		return err
	}

	packageName := c.Args().Get(0)
	if packageName != "" {
		for _, file := range config.Source {
			if file.Name == packageName {
				filePath := module.DownloadFileToCache(file.URL)
				err = installPackage(filePath)
				if err != nil {
					return err
				}
				return nil
			}
		}
		color.Redln("Package not found.")
		return errors.New("1")
	} else {
		color.Redln("No package specified.")
		return errors.New("1")
	}
}

func installPackage(filePath string) error {
	fileName := filepath.Base(filePath)

	color.Yellowln("Installing", fileName)
	module.FullWidthMessage("installation log start", color.Gray)
	err := module.InstallPackageWithFilePath(filePath)
	if err != nil {
		return err
	}

	module.FullWidthMessage("installation log end", color.Gray)
	color.Greenln("Successfully installed", fileName)
	return nil
}
