package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// Config is the settings.ini representation
type Config struct {
	DeviceID string
	Broker   string
	Username string
	Password string
	URL      string
	Token    string
	Topic    string
}

func getINIValue(cfgInstance *ini.File, section string, key string) string {
	value := cfgInstance.Section(section).Key(key).String()
	if value == "" {
		fmt.Printf("The field %v cann't be empty", key)
		os.Exit(1)
	}
	return value
}

func (config *Config) fill(cfgInstance *ini.File) {
	config.DeviceID = getINIValue(cfgInstance, "device", "device_id")
	config.Broker = getINIValue(cfgInstance, "mqtt", "broker")
	config.Topic = getINIValue(cfgInstance, "mqtt", "topic")
	config.Username = getINIValue(cfgInstance, "mqtt", "username")
	config.Password = getINIValue(cfgInstance, "mqtt", "password")
	config.URL = getINIValue(cfgInstance, "api", "url")
	config.Token = getINIValue(cfgInstance, "api", "token")
}

// LoadConfig returns a Config struct
func LoadConfig(path string) Config {
	cfgInstance := loadFile(path)

	var cfg Config
	cfg.fill(cfgInstance)
	return cfg
}
