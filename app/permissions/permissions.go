package app

import (
	"encoding/json"
	"os"
)

type Config struct {
	UserRoles map[string]UserRole `json:"user_roles"`
}

type UserRole struct {
	Permissions []string `json:"permissions"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("app/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func CheckPermission(userRole string, permission string, config *Config) bool {
	role, exists := config.UserRoles[userRole]
	if !exists {
		return false
	}

	for _, perm := range role.Permissions {
		if perm == permission {
			return true
		}
	}
	return false
}
