package commands

import (
	"log"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(c *cli.Context) {
	module.CreateCacheDotFolder()
	config, err := module.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	packageName := c.String("name")
	if packageName != "" {
		for _, file := range config.Files {
			if file.Name == packageName {
				filePath := module.DownloadFileToCache(file.URL)
				installPackage(filePath)
				return
			}
		}
	} else {
		log.Fatal("Package not found")
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
