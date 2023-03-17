package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DB struct {
		Url         string `yaml:"url"`
		MaxOpenConn int    `yaml:"max_open_conn"`
		MaxIdleConn int    `yaml:"max_idle_conn"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func NewConfig(path string) (*Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	instance := &Config{}
	err = yaml.Unmarshal(yamlFile, &instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
