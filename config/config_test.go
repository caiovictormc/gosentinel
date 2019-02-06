package config_test

import (
	"fmt"
	"sync"
	"testing"

	. "github.com/caiovictormc/goteway/config"
)

const (
	_CONFIG_INI_FILE = "testdata/settings.test.ini"
	failMark         = "\u2717"
	checkMark        = "\u2713"
)

func checkField(
	value string, key string, expected string,
	status chan map[int]string, wait *sync.WaitGroup,
) {
	if value != expected {
		msg := fmt.Sprintf("cfg.%s cann't be %s / Expected: %s", key, value, expected)
		status <- map[int]string{1: msg}
	}
	wait.Done()
}

func TestLoadConfig(t *testing.T) {
	var waitGroup sync.WaitGroup
	status := make(chan map[int]string, 7)

	cfg := LoadConfig(_CONFIG_INI_FILE)

	waitGroup.Add(7)

	go checkField(cfg.DeviceID, "device_id", "awesome-uuid", status, &waitGroup)
	go checkField(cfg.Broker, "broker", "localhost:1883", status, &waitGroup)
	go checkField(cfg.Username, "username", "awesome-user-device", status, &waitGroup)
	go checkField(cfg.Password, "password", "awesome-password-device", status, &waitGroup)
	go checkField(cfg.Topic, "main_topic", "awesome-topic", status, &waitGroup)
	go checkField(cfg.URL, "url", "localhost:8000", status, &waitGroup)
	go checkField(cfg.Token, "token", "awesome-token", status, &waitGroup)

	waitGroup.Wait()
	close(status)
	for logs := range status {
		for log := range logs {
			if log == 1 {
				t.Fatalf(logs[log])
			}
		}
	}
}
