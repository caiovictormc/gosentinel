package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// Config is the settings.ini representation
type Config struct {
	deviceID string
	broker   string
	username string
	password string
	url      string
	token    string
}

func (config *Config) fill(cfgInstance *ini.File, structure map[string][]string) {
	for section, fields := range structure {
		for _, field := range fields {
			if cfgInstance.Section(section).Key(field).String() == "" {
				fmt.Printf("The field %v cann't be empty", field)
				os.Exit(1)
			}
		}
	}
}

// LoadConfig returns a Config struct
func LoadConfig(path string) Config {
	structure := map[string][]string{
		"device": []string{"device_id"},
		"mqtt":   []string{"broker", "username", "password"},
		"api":    []string{"url", "token"},
	}

	cfgInstance := loadFile(path)

	var cfg Config
	fmt.Println(cfg)
	cfg.fill(cfgInstance, structure)
	fmt.Println(cfg)
	return cfg
}
