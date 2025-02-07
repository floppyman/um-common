package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/umbrella-sh/um-common/ext"
	"github.com/umbrella-sh/um-common/logging/ulog"
)

// (?m) is the go way of specifying flags for the regex compiler
const rxStr = `(?m)//.*$`

var re = regexp.MustCompile(rxStr)

//goland:noinspection GoUnusedExportedFunction
func LoadConfig(path string, cfgObjInst any, hasJsonComments bool, doLogging bool) error {
	if doLogging {
		ulog.Console.Info().Msg("Loading config...")
	}

	configPath := filepath.Join(ext.WorkingDir(), path)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("could not open config file, ERR: %s", err)
	}

	var newData []byte
	if hasJsonComments {
		newData = jsonCommentRemover(data)
	}

	err = json.Unmarshal(newData, &cfgObjInst)
	if err != nil {
		return fmt.Errorf("could not parse config file, ERR: %s", err)
	}

	if doLogging {
		ulog.Console.Info().Msg("Done")
	}
	return nil
}

func jsonCommentRemover(data []byte) []byte {
	newData := re.ReplaceAllString(string(data), "")
	return []byte(newData)
}
