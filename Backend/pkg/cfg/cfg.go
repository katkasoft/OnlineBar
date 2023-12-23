package cfg

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Database_Config struct {
	Database struct {
		DBUser string `yaml:"db_user"`
		DBPass string `yaml:"db_pass"`
		DBNet  string `yaml:"db_net"`
		DBAddr string `yaml:"db_addr"`
		DBName string `yaml:"db_name"`
		DBPort string `yaml:"db_port"`
	} `yaml:"database"`
}

type Server_Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func DBConfig() Database_Config {
	data, err := os.ReadFile("configs/db.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config Database_Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	return config
}

func ServerConfig() Server_Config {
	data, err := os.ReadFile("configs/server_config.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config Server_Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	return config
}
