package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	RTMP struct {
		Port string `yaml:"port"`
	} `yaml:"rtmp"`
}

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return &cfg
}