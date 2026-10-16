package appconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	DataBase struct {
		Sqlite struct {
			Address	   	     string `yaml:"address"`
		} `yaml:"sqlite"`
	} `yaml:"database"`
	Secret string `yaml:"secret"`
}

func NewConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := new(Config)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
