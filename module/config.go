package module

import (
	"encoding/json"
	"fmt"
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
	Files []FileLink `json:"files"`
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

func SetFileURLConfigs(name string, URL string) (ConfigFileJSON, error) {
	config, err := GetConfigs()
	if err != nil {
		return config, err
	}

	for i, file := range config.Files {
		if file.Name == name {
			fmt.Printf("The URL of the name is already exist. Would you like to change the URL of %s? (y/n) ", name)
			var answer string
			fmt.Scanln(&answer)

			if answer == "y" || answer == "Y" {
				config.Files[i].URL = URL
				return config, nil
			} else {
				fmt.Println("No changes were made")
				return config, nil
			}
		}
	}

	// append new file
	config.Files = append(config.Files, FileLink{Name: name, URL: URL})

	return config, nil
}

func SaveToConfigFile(config ConfigFileJSON) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configFilePath := filepath.Join(home, ".ui", "config.json")

	if err := os.Remove(configFilePath); err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(configBytes)
	if err != nil {
		return err
	}

	return nil
}
