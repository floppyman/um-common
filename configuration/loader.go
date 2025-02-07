package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

// LoadJson loads configuration from json file, the config object should define both 'json' and 'envconfig' field names
//
//goland:noinspection GoUnusedExportedFunction
func LoadJson[TConfig any](filePath string, cfg *TConfig) (*TConfig, error) {
	return load("json", filePath, cfg)
}

// LoadYaml loads configuration from yaml file, the config object should define both 'yaml' and 'envconfig' field names
//
//goland:noinspection GoUnusedExportedFunction
func LoadYaml[TConfig any](filePath string, cfg *TConfig) (*TConfig, error) {
	return load("yaml", filePath, cfg)
}

func load[TConfig any](configType string, filePath string, cfg *TConfig) (*TConfig, error) {
	if filePath == "" {
		return cfg, fmt.Errorf("empty configuration file path")
	}

	if configType == "json" {
		_ = readJsonFile(filePath, cfg)
	} else if configType == "yaml" {
		_ = readYamlFile(filePath, cfg)
	}

	err := readEnv(cfg)
	return cfg, err
}

func readYamlFile[TConfig any](filePath string, cfg *TConfig) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func readJsonFile[TConfig any](filePath string, cfg *TConfig) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func readEnv[TConfig any](cfg *TConfig) (err error) {
	err = envconfig.Process("", cfg)
	return err
}
