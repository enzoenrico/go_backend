package app

import (
	"encoding/json"
	"os"
)

type Config struct {
	JwtSecret    string `json:"jwt_secret"`
	DatabaseFile string `json:"database_file"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("app/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
