package main

import (
	"flag"
	"fmt"

	"github.com/galexrt/desktop-helper/pkg/config"
)

var (
	configFilename string
)

func init() {
	flag.StringVar(&configFilename, "config", "./config.example.yaml", "Config file location")
}

func main() {
	flag.Parse()
	cfg, _ := config.Read(configFilename)
	fmt.Printf("%+v", cfg)
}
