package commands

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(c *cli.Context) {
	var cmd = exec.Command("sudo", "-v")
	cmd.Stdout = nil
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	module.CreateCacheDotFolder()
	config, err := module.GetSource()
	if err != nil {
		log.Fatal(err)
	}

	packageName := c.Args().Get(0)
	if packageName != "" {
		for _, file := range config.Source {
			if file.Name == packageName {
				filePath := module.DownloadFileToCache(file.URL)
				installPackage(filePath)
				return
			}
		}
		log.Fatal("Package not found")
	}

	for _, file := range config.Source {
		filePath := module.DownloadFileToCache(file.URL)
		installPackage(filePath)
	}
}

func installPackage(filePath string) {
	fileName := filepath.Base(filePath)

	color.Yellowln("Installing", fileName)
	module.FullWidthMessage("installation log start", color.Gray)
	err := module.InstallPackageWithFilePath(filePath)
	if err != nil {
		log.Fatal(err)
	}

	module.FullWidthMessage("installation log end", color.Gray)
	color.Greenln("Successfully installed", fileName)
}
