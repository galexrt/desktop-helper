package main

import (
	"flag"
	"log"
	"sync"

	"github.com/galexrt/desktop-helper/detector"
	"github.com/galexrt/desktop-helper/pkg/config"
)

var (
	configFilename string
)

func init() {
	flag.StringVar(&configFilename, "config", "./config.yaml", "Config file location")
}

func main() {
	flag.Parse()
	cfg, _ := config.Read(configFilename)
	detect, err := detector.NewDetector(cfg)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		detect.Run()
		wg.Done()
	}()
	wg.Wait()
}
