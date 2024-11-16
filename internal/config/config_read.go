package config

import (
	"encoding/json"
	"log"
	"os"
)

func Read() config {
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("cant get config file path: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var cfg config
	err = json.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("cant Unmarshal json: %v", err)
	}
	return cfg
}
