package ext

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

//goland:noinspection GoUnusedExportedFunction
func WorkingDir() string {
	if os.Getenv("ENV") == "debug" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return wd
	} else {
		ep, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		epd := filepath.Dir(ep)
		return epd
	}
}

//goland:noinspection GoUnusedExportedFunction
func FileExist(path string) bool {
	_, err := os.OpenFile(path, os.O_RDONLY, 0)
	return !errors.Is(err, os.ErrNotExist)
}
