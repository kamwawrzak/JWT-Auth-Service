package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerCfg ServerCfg `yaml:"server"`
	DbCfg DbCfg `yaml:"database"`
	JWTCfg JWTCfg `yaml:"jwt"`
}

type ServerCfg struct {
	Port int `yaml:"port"`
}

type DbCfg struct {
	Reader DsnCfg `yaml:"reader"`
	Writer DsnCfg `yaml:"writer"`
	Connections ConnsCfg `yaml:"connections"`
}

type ConnsCfg struct {
	MaxConnLifetime time.Duration `yaml:"maxConnLifetimeInMinutes"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxIdleConns int `yaml:"maxIdleConns"`
}

type DsnCfg struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JWTCfg struct {
	SecretKey string `yaml:"secret_key"`
	TimeToLive time.Duration `yaml:"time_to_live"`
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
