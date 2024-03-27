package module

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func CreateCacheDotFolder() {
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

func DownloadFileToCache(url string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cachePath := filepath.Join(home, ".ui", "cache")
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		err = os.Mkdir(cachePath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(filepath.Join(cachePath, filepath.Base(url)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return file.Name()

}
