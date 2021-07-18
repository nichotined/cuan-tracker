package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		DBName  string `yaml:"database_name"`
		Address string `yaml:"address"`
	} `yaml:"database"`
}

func GetConfig(configPath string) Config {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
