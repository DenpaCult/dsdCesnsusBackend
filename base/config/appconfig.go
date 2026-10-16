package appconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
		TLS  bool   `yaml:"tls"`
		MessageLimit int `yaml:"message_limit"`
	} `yaml:"server"`

	DataBase struct {
		Sqlite struct {
			Address	   	     string `yaml:"address"`
		} `yaml:"sqlite"`
	} `yaml:"database"`
	Secret string `yaml:"secret"`
	S3 struct {
		Region    string `yaml:"region"`
		Bucket    string `yaml:"bucket"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
		Endpoint  string `yaml:"endpoint"`
	} `yaml:"s3"`
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
