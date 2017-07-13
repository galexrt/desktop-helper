package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/galexrt/desktop-helper/detector"
	"github.com/galexrt/desktop-helper/pkg/config"
	"github.com/prometheus/common/log"

	// Triggers need to be imported for their init() to becalled and register themselves
	_ "github.com/galexrt/desktop-helper/pkg/triggers/acpid"
	_ "github.com/galexrt/desktop-helper/pkg/triggers/hosts"
	_ "github.com/galexrt/desktop-helper/pkg/triggers/network"
	_ "github.com/galexrt/desktop-helper/pkg/triggers/screens"

	// Actions need to be imported for their init() to becalled and register themselves
	_ "github.com/galexrt/desktop-helper/pkg/actions/exec"
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
		log.Error(detect.Run(ctx))
		cancel()
	}()
	wg.Wait()
	if sig.String() == "" {
		main()
	}
}
