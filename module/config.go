package module

import (
	"log"
	"os"
	"path/filepath"
)

func CreateConfigFile() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// create dot folder if it is not exist.
	CreateCacheDotFolder()

	path := filepath.Join(home, ".ui")
	configPath := filepath.Join(path, "config.json")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString("{}")
		if err != nil {
			log.Fatal(err)
		}
	}
}
