package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func (cfg config) SetUser(current_user_name string) {
	cfg.CurrentUserName = current_user_name
	write(cfg)
}

func getConfigFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, "gatorconfig.json")
	return path, nil
}

func write(cfg config) error {
	path, err := getConfigFilePath()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
