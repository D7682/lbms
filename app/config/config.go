package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DSN  string `json:"dsn" yaml:"dsn"`
	Port string `json:"port" yaml:"port"`
}

func New(fileName string) (*Config, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
