package config

import (
	"encoding/json"
	"io/ioutil"
)

const FileName = ".maestro.json"

type Config struct {
	Aliases map[string]Alias
}

func Load(path string) (*Config, error) {
	fil, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(fil, &config); err != nil {
		return nil, err
	}

	// Ensure the Alias names are set on the alias structs
	for name, alias := range config.Aliases {
		alias.Name = name
		config.Aliases[name] = alias
	}

	return &config, nil
}
