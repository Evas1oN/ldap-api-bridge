package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type KlambriConfig struct {
	Connector struct {
		Realm    string `yaml:"realm"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Secure   bool   `yaml:"secure"`
	} `yaml:"connector"`

	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func ReadConfig(path string) (*KlambriConfig, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	println(string(buffer[:]))

	config := &KlambriConfig{}

	if err := yaml.Unmarshal(buffer, config); err != nil {
		return nil, err
	}

	return config, nil
}
