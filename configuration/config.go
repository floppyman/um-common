package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type FuncUnmarshal func(in []byte, out interface{}) (err error)

// (?m) is the go way of specifying flags for the regex compiler
const rxStr = `(?m)(//.*|#.*)$`

var re = regexp.MustCompile(rxStr)

// LoadJson reads the json file from filePath, and overrides fields in config based on passed environment variables with optional prefix.
//
// Comments are removed before loading if hasComments is true, if there are comments and this is set to false, loading will fail.
//
// The config object should define both 'json' and 'envconfig' field names, for consistence and so that it can be saved again with SaveJson.
func LoadJson[TConfig any](filePath string, defaultConfig *TConfig, hasComments bool, envPrefixName string) (*TConfig, error) {
	return load(json.Unmarshal, filePath, defaultConfig, hasComments, envPrefixName)
}

// LoadYaml reads the json file from filePath, and overrides fields in config based on passed environment variables with optional prefix.
//
// Comments are removed before loading if hasComments is true, if there are comments and this is set to false, loading will fail.
//
// The config object should define both 'json' and 'envconfig' field names, for consistence and so that it can be saved again with SaveYaml.
func LoadYaml[TConfig any](filePath string, defaultConfig *TConfig, hasComments bool, envPrefixName string) (*TConfig, error) {
	return load(yaml.Unmarshal, filePath, defaultConfig, hasComments, envPrefixName)
}

func load[TConfig any](unmarshal FuncUnmarshal, filePath string, config *TConfig, hasComments bool, envPrefixName string) (*TConfig, error) {
	if filePath == "" {
		return config, fmt.Errorf("configuration file path is empty, must have a value")
	}

	fc, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var newFc = fc
	if hasComments {
		newFc = commentRemover(fc)
	}

	err = unmarshal(newFc, config)
	if err != nil {
		return nil, err
	}

	err = readEnv(config, envPrefixName)
	return config, err
}

func readEnv[TConfig any](config *TConfig, envPrefixName string) (err error) {
	err = envconfig.Process(envPrefixName, config)
	return err
}

func commentRemover(data []byte) []byte {
	newData := re.ReplaceAllString(string(data), "")
	return []byte(newData)
}

type FuncMarshal func(in interface{}) (out []byte, err error)

// SaveJson serializes the config to JSON and creates a file at filePath.
//
// If shouldOverride is false, and a filePath exists, then the config file will not be written.
//
// The config object should define both 'json' and 'envconfig' field names, so that it can be easily loaded with LoadJson.
//
// SaveJson returns 'true, nil' when the file is saved, or shouldOverride is false and file exists, else it returns 'false, err'
func SaveJson(filePath string, config interface{}, shouldOverride bool) error {
	return save(json.Marshal, filePath, config, shouldOverride)
}

// SaveYaml serializes the config to YAML and creates a file at filePath.
//
// If shouldOverride is false, and a filePath exists, then the config file will not be written.
//
// The config object should define both 'yaml' and 'envconfig' field names, so that it can be easily loaded with LoadYaml.
//
// SaveYaml returns 'true, nil' when the file is saved, or shouldOverride is false and file exists, else it returns 'false, err'
func SaveYaml(filePath string, config interface{}, shouldOverride bool) error {
	return save(yaml.Marshal, filePath, config, shouldOverride)
}

func save(marshal FuncMarshal, filePath string, config interface{}, shouldOverride bool) error {
	if filePath == "" {
		return fmt.Errorf("configuration file path is empty, must have a value")
	}

	if config == nil || reflect.ValueOf(config).Kind() != reflect.Struct {
		return fmt.Errorf("configuration is nil or must be a struct")
	}

	if !shouldOverride {
		return nil
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrExist) {
		err = os.Remove(filePath)
		if err != nil {
			return fmt.Errorf("failed to remove existing config file at: %s", filePath)
		}
	}

	bytes, err := marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
