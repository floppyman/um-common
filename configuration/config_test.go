package configuration

import (
	"os"
	"testing"
)

type saverTestConfig struct {
	Name       string `json:"name" yaml:"name"`
	Address    string `json:"address" yaml:"address"`
	PostalCode int    `json:"postal_code" yaml:"postal_code"`
	IsBusiness bool   `json:"is_business" yaml:"is_business"`
}

func cleanUp(jsonYaml string) {
	err := os.Remove("./test." + jsonYaml)
	if err != nil {
		return
	}
}

func TestSaveJson(t *testing.T) {
	cfg := saverTestConfig{
		Name:       "test",
		Address:    "street 19",
		PostalCode: 1234,
		IsBusiness: true,
	}
	err := SaveJson("test.json", cfg, true)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveJsonIndented(t *testing.T) {
	cfg := saverTestConfig{
		Name:       "test",
		Address:    "street 19",
		PostalCode: 1234,
		IsBusiness: true,
	}
	err := SaveJsonIndented("test.json", cfg, true)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveYaml(t *testing.T) {
	cfg := saverTestConfig{
		Name:       "test",
		Address:    "street 19",
		PostalCode: 1234,
		IsBusiness: true,
	}
	err := SaveYaml("test.yaml", cfg, true)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadJson(t *testing.T) {
	var cfgDef saverTestConfig
	_, err := LoadJson("test.json", &cfgDef, false, "")
	if err != nil {
		t.Error(err)
	}
	cleanUp("json")
}

func TestLoadYaml(t *testing.T) {
	var cfgDef saverTestConfig
	_, err := LoadYaml("test.yaml", &cfgDef, false, "")
	if err != nil {
		t.Error(err)
	}
	cleanUp("yaml")
}
