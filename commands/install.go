package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(*cli.Context) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	module.CreateCacheDotFolder()
	config, err := module.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range config.Files {
		fmt.Println(file.Name + " Downloading... " + " " + file.URL)
		module.DownloadFileToCache(file.URL)
	}

	fmt.Println("Download complete!")

	for _, file := range config.Files {
		fmt.Println("Installing... " + file.Name)
		err := module.InstallPackageWithFilePath(home + "/.ui/cache/" + filepath.Base(file.URL))
		if err != nil {
			log.Fatal(err)
		}
	}
}
