package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"

	yaml "gopkg.in/yaml.v2"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/poller"
	log "github.com/sirupsen/logrus"

	// Triggers
	_ "github.com/galexrt/desktop-helper/poller/triggers/ipaddress"
	_ "github.com/galexrt/desktop-helper/poller/triggers/xrandr"
	// Actions
	_ "github.com/galexrt/desktop-helper/runner/actions/exec"
	_ "github.com/galexrt/desktop-helper/runner/actions/libnotify"
)

var (
	configFilename string
	debug          bool
)

func init() {
	flag.StringVar(&configFilename, "config", "./config.yaml", "Config file location")
	flag.BoolVar(&debug, "debug", true, "Enable debug mode")
}

func main() {
	flag.Parse()
	cfg, err := config.Read(configFilename)
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		log.SetLevel(log.DebugLevel)
		printed, _ := yaml.Marshal(cfg)
		fmt.Printf("%+v\n", string(printed))
	}

	poller, err := poller.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var sig os.Signal
	go func() {
		select {
		case sig = <-c:
			log.Infof("Signal received: %s", sig)
		case <-ctx.Done():
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Error(poller.Run(ctx))
		cancel()
	}()
	wg.Wait()
	if sig != nil && sig.String() == "" {
		main()
	}
}
