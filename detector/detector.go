package detector

import (
	"context"
	"strconv"
	"time"

	"github.com/galexrt/desktop-helper/pkg/config"
	"github.com/galexrt/desktop-helper/pkg/triggers"
	"github.com/prometheus/common/log"
)

type Detector struct {
	activeTriggers map[string]interface{}
	checkInterval  time.Duration
	timeout        time.Duration
}

func NewDetector(config *config.Config) (*Detector, error) {
	detect := &Detector{
		activeTriggers: make(map[string]interface{}),
	}
	for profileName, profile := range config.Profiles {
		for name, triggerConf := range profile.Trigger {
			if triggers.Exist(name) {
				log.With("module", "detector").With("func", "NewDetector").
					Fatalf("Given trigger '%s' in profile '%s' does not exist.\n", name, profileName)
			}
			if _, ok := detect.activeTriggers[name]; !ok {
				detect.activeTriggers[name] = triggerConf
				log.With("module", "detector").With("func", "NewDetector").
					Debugf("Trigger added to check list: '%s'\n", name)
			}
		}
	}

	var err error
	duration := strconv.FormatInt(config.DetectorOptions.PollInterval, 10)
	if detect.checkInterval, err = time.ParseDuration(duration + "s"); err != nil {
		return detect, err
	}
	log.With("module", "detector").With("func", "NewDetector").
		Debugf("config.Options.PollInterval: '%s'\n", detect.checkInterval)
	timeout := strconv.FormatInt(config.DetectorOptions.PollInterval, 10)
	if detect.timeout, err = time.ParseDuration(timeout + "s"); err != nil {
		return detect, err
	}
	log.With("module", "detector").With("func", "NewDetector").
		Debugf("config.Options.Timeout: '%s'\n", detect.timeout)
	return detect, nil
}

func (detect Detector) Run() {
	errs := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), detect.timeout)
	// TODO Implement check and evaluate logic
	// 1. Every PollInterval loop over triggers to get the latest state `GetState()`
	// 2. Evaluate after the loop succeded withtout errors
	// 2.1. Loop over every profile's conditions
	// 2.1.1. Match with current states
	// 2.1.2. If there is a full match, go to exec of on_enable, set state in
	//      global Detector var, return immediately.
	// 3. Go to step 1
	var err error
	go func() {
		for {
			err = <-errs
			log.Errorln(err)
			cancel()
		}
	}()
	for {
		log.With("module", "detector").Debug("Triggers checking ...")
		go func(ctx context.Context) {

		}(ctx)
		log.With("module", "detector").Debug("Triggers checked.")
		<-time.After(detect.checkInterval)
	}
}
