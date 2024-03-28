package commands

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func Install(*cli.Context) {
	module.CreateCacheDotFolder()
	config, err := module.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range config.Files {
		fmt.Println(file.Name + " Downloading... " + " " + file.URL)
		module.DownloadFileToCache(file.URL)
	}
}
