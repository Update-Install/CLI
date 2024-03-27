package module

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type FileLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ConfigFileJSON struct {
	Files FileLink
}

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

func GetConfigs() (ConfigFileJSON, error) {
	var config ConfigFileJSON

	home, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	CreateConfigFile()

	configFilePath := filepath.Join(home, ".ui", "config.json")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return config, err
	}

	// Read the config file
	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(configFileBytes, &config); err != nil {
		return config, err
	}

	return config, nil
}
