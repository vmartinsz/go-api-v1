package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	User     string `yaml:"db_user"`
	Port     int    `yaml:"db_port"`
	Password string `yaml:"db_password"`
	Host     string `yaml:"db_host"`
	Name     string `yaml:"db_name"`
}

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var conf Config

	err = yaml.Unmarshal(data, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

var config *Config

func Get() *Config {
	if config == nil {
		config, err := Load("../../../config/config.yml")
		if err != nil {
			log.Println("não foi possível carregar este arquivo")
		}
		return config
	}

	return config
}
