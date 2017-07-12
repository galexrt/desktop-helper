package detector

import (
	"context"
	"reflect"
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
	triggersCache  map[string]triggers.Trigger
	profiles       map[string]interface{}
}

// NewDetector creates a new detector
func NewDetector(config *config.Config) (*Detector, error) {
	detect := &Detector{
		activeTriggers: make(map[string]interface{}),
		triggersCache:  make(map[string]triggers.Trigger),
		profiles:       make(map[string]interface{}),
	}
	for profileName, profile := range config.Profiles {
		for name, triggerConf := range profile.Trigger {
			detect.profiles[name] = triggerConf
			if !triggers.Exist(name) {
				log.With("module", "detector").With("func", "NewDetector").
					Fatalf("Given trigger '%s' in profile '%s' does not exist.\n", name, profileName)
			}
			if _, ok := detect.activeTriggers[name]; !ok {
				detect.activeTriggers[name] = struct{}{}
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
func (detect Detector) Run(bctx context.Context) error {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(bctx)
	wg.Add(1)
	errors := make(chan error)
	go func() {
		defer wg.Done()
		for {
			log.With("module", "detector").Debug("Triggers checking ...")
			states, err := detect.getStateFromTriggers(ctx)
			if err != nil {
				errors <- err
				return
			}
			err = detect.evaluateStateFromTriggers(states)
			if err != nil {
				errors <- err
				return
			}
			log.With("module", "detector").Debug("Triggers checked.")
			select {
			case <-ctx.Done():
				return
			case <-time.After(detect.checkInterval):
			}
		}
	}()

	var err error
	err = nil
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errors:
	}
	cancel()
	wg.Wait()
	return err
}

func (detect Detector) getStateFromTriggers(ctx context.Context) (map[string]interface{}, error) {
	wg := sync.WaitGroup{}
	tctx, cancel := context.WithCancel(ctx)
	errors := make(chan error)
	results := make(map[string]interface{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		var mutex = &sync.Mutex{}
		// Get the current state of all activeTriggers
		for triggerName := range detect.activeTriggers {
			var err error
			if _, ok := detect.triggersCache[triggerName]; !ok {
				detect.triggersCache[triggerName] = triggers.Get(triggerName)
			}
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				triggerConf := detect.activeTriggers[name]
				mutex.Lock()
				results[name], err = detect.triggersCache[name].GetState(ctx, triggerConf)
				mutex.Unlock()
				if err != nil {
					errors <- err
				}
			}(triggerName)
		}
	}()
	wgc := make(chan struct{})
	go func() {
		defer close(wgc)
		wg.Wait()
	}()
	var err error
	select {
	case err = <-errors:
	case <-wgc:
	case <-tctx.Done():
	case <-ctx.Done():
	}
	cancel()

	return results, err
}

func (detect Detector) evaluateStateFromTriggers(states map[string]interface{}) error {
	for profile, conf := range detect.profiles {
		if _, ok := states[profile]; !ok {
			continue
		}
		log.Debugf("state: %+v\n", states[profile])
		log.Debugf("profile: %+v\n", conf)
		log.Debugf("state == profile: %+v\n", reflect.DeepEqual(conf, states[profile]))
		if reflect.DeepEqual(conf, states[profile]) {
			log.Debugf("Profile '%s' matched with current state.\n", profile)
			return nil
		}
	}
	return nil
}
