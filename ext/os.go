package ext

import (
	"log"
	"os"
	"path/filepath"
)

// WorkingDir tries to get the correct path of the current execution based on the value of the ENV environment variable.
// if the ENV is debug, then it will use os.Getwd() else it will use filepath.Dir(os.Executable())
//
//goland:noinspection GoUnusedExportedFunction
func WorkingDir() string {
	if os.Getenv("ENV") == "debug" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return wd
	}

	ep, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	epd := filepath.Dir(ep)
	return epd
}

// FileExist checks if the provided `filePath` is an existing file on disk
//
//goland:noinspection GoUnusedExportedFunction
func FileExist(filePath string) bool {
	return diskEntityExists(filePath)
}

// FolderExist checks if the provided `folderPath` is an existing folder on disk
//
//goland:noinspection GoUnusedExportedFunction
func FolderExist(folderPath string) bool {
	return diskEntityExists(folderPath)
}

func diskEntityExists(entityPath string) bool {
	if p, _ := filepath.Glob(entityPath); len(p) > 0 {
		return true
	}
	return false
}
