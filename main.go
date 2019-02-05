package main

import (
	"fmt"

	conf "github.com/caiovictormc/gosentinel/config"
)

func main() {
	cfg := conf.LoadConfig("settings.ini")
	fmt.Println(cfg)
}
