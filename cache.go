package main

import (
	"log"
	"os"
	"path/filepath"
)

func createCacheDotFolder() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(home, ".ui")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
