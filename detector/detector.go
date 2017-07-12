package detector

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/galexrt/desktop-helper/pkg/config"
	"github.com/galexrt/desktop-helper/pkg/triggers"
	"github.com/prometheus/common/log"
)

// Detector containing configuration for the detector itself
type Detector struct {
	activeTriggers map[string]interface{}
	checkInterval  time.Duration
	timeout        time.Duration
}

// NewDetector creates a new detector
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
		Debugf("config.DetectorOptions.PollInterval: '%s'\n", detect.checkInterval)
	timeout := strconv.FormatInt(config.DetectorOptions.PollInterval, 10)
	if detect.timeout, err = time.ParseDuration(timeout + "s"); err != nil {
		return detect, err
	}
	log.With("module", "detector").With("func", "NewDetector").
		Debugf("config.DetectorOptions.Timeout: '%s'\n", detect.timeout)
	return detect, nil
}

// Run run the detector logic
func (detect Detector) Run(ctx context.Context) error {
	// TODO Implement check and evaluate logic
	// 1. DONE - Every PollInterval loop over triggers to get the latest state `GetState()`
	// 2. WIP - Evaluate after the loop succeded withtout errors
	// 2.1. WIP - Loop over every profile's conditions
	// 2.1.1. WIP - Match with current states
	// 2.1.2. WIP - If there is a full match, go to exec of on_enable, set state in
	//      global Detector var, return immediately.
	// 3. WIP - Go to step 1
	wg := sync.WaitGroup{}
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		default:
		}
		log.With("module", "detector").Debug("Triggers checking ...")
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			// Get the current state
			// if there is an error
			select {
			case <-ctx.Done():
				return
			default:
			}
			// Evaluate the state
		}(ctx)
		log.With("module", "detector").Debug("Triggers checked.")
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		case <-time.After(detect.checkInterval):
		}
	}
}
