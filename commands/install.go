package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(c *cli.Context) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	module.CreateCacheDotFolder()
	config, err := module.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	packageName := c.String("name")
	if packageName != "" {
		for _, file := range config.Files {
			if file.Name == packageName {
				installPackage(home, file.URL)
				return
			}
		}
		log.Fatal("Package not found")
	}

	for _, file := range config.Files {
		fmt.Println(file.Name + " Downloading... " + file.URL)
		module.DownloadFileToCache(file.URL)
	}

	fmt.Println("Download complete!")
	for _, file := range config.Files {
		installPackage(home, file.URL)
	}
}

func installPackage(home, filePath string) {
	fileName := filepath.Base(filePath)
	err := module.InstallPackageWithFilePath(home + "/.ui/cache/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
}
