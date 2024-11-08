package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(name string) (*ConfigDB, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	cfg := &ConfigDB{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
