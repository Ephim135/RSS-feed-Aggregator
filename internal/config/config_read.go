package config

import (
	"encoding/json"
	"log"
	"os"
)

func Read() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("cant get config file path: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var cfg *Config
	err = json.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("cant Unmarshal json: %v", err)
	}
	return cfg, nil
}
