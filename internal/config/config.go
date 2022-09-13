package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HTTP *HTTP  `yaml:"http"`
	Mode string `yaml:"mode"`
}

type HTTP struct {
	Port   string `yaml:"port"`
	Domain string `yaml:"domain"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	d := yaml.NewDecoder(f)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
