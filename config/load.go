package config

import (
	"fmt"
	"os"

	ini "gopkg.in/ini.v1"
)

func loadFile(path string) *ini.File {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read the config file: %v", err)
		os.Exit(1)
	}

	return cfg
}
