package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerCfg ServerCfg `yaml:"server"`
}

type ServerCfg struct {
	Port int `yaml:"port"`
}

func NewConfig(configPath string) (*Config, error) {
    config := &Config{}
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    d := yaml.NewDecoder(file)

    // Start YAML decoding from file
    if err := d.Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}