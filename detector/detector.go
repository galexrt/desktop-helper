package detector

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	cfg "github.com/galexrt/desktop-helper/pkg/config"
	"github.com/galexrt/desktop-helper/pkg/triggers"
	"github.com/galexrt/desktop-helper/runner"
	"github.com/prometheus/common/log"
)

// Detector containing configuration for the detector itself
type Detector struct {
	activeTriggers map[string]map[string]interface{}
	checkInterval  time.Duration
	timeout        time.Duration
	triggersCache  map[string]triggers.Trigger
	profiles       map[string]cfg.Profile
	currentProfile string
	runner         *runner.Runner
}

func getTriggerName(name string) string {
	return strings.SplitN(name, "_", 2)[0]
}

// NewDetector creates a new detector
func NewDetector(config *cfg.Config) (*Detector, error) {
	detect := &Detector{
		activeTriggers: make(map[string]map[string]interface{}),
		triggersCache:  make(map[string]triggers.Trigger),
		profiles:       make(map[string]cfg.Profile),
		runner:         runner.New(config),
	}
	detect.profiles = config.Profiles

	for profileName, profile := range config.Profiles {
		for realName, triggerConf := range profile.Trigger {

			if !triggers.Exist(getTriggerName(realName)) {
				return detect, fmt.Errorf("given trigger '%s' in profile '%s' does not exist", getTriggerName(realName), profileName)
			}
			if _, ok := detect.activeTriggers[realName]; !ok {
				detect.activeTriggers[realName] = triggerConf
				log.With("func", "NewDetector").Debugf("Trigger added to check list: '%s'", realName)
			}
		}
	}

	var err error
	duration := strconv.FormatInt(config.DetectorOptions.PollInterval, 10)
	if detect.checkInterval, err = time.ParseDuration(duration + "s"); err != nil {
		return detect, err
	}
	log.With("func", "NewDetector").Debugf("config.DetectorOptions.PollInterval: '%s'", detect.checkInterval)
	timeout := strconv.FormatInt(config.DetectorOptions.PollInterval, 10)
	if detect.timeout, err = time.ParseDuration(timeout + "s"); err != nil {
		return detect, err
	}
	log.With("func", "NewDetector").Debugf("config.DetectorOptions.Timeout: '%s'", detect.timeout)
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
			log.Debug("Triggers running ...")
			states, err := detect.getStateFromTriggers(ctx)
			if err != nil {
				errors <- err
				return
			}
			profile, match := detect.evaluateStateFromTriggers(states)
			log.Debug("Triggers ran and evaluated.")
			if match {
				if detect.currentProfile == "" {
					detect.currentProfile = profile
				}
				if detect.currentProfile != profile {
					err = detect.runner.OnDisable(ctx, detect.currentProfile)
					if err != nil {
						errors <- err
						return
					}
					detect.currentProfile = profile
				}
				err = detect.runner.OnEnable(ctx, detect.currentProfile)
				if err != nil {
					errors <- err
					return
				}
			}
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
	results := map[string]interface{}{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var mutex = &sync.Mutex{}
		// Get the current state of all activeTriggers
		for triggerName := range detect.activeTriggers {
			if _, ok := detect.triggersCache[triggerName]; !ok {
				detect.triggersCache[getTriggerName(triggerName)] = triggers.Get(getTriggerName(triggerName))
			}
			wg.Add(1)
			go func(name string, trigger triggers.Trigger, triggerConf map[string]interface{}) {
				defer wg.Done()
				result, err := trigger.GetState(ctx, name, triggerConf)
				mutex.Lock()
				results[name] = result
				mutex.Unlock()
				if err != nil {
					errors <- err
				}
			}(triggerName, detect.triggersCache[getTriggerName(triggerName)], detect.activeTriggers[triggerName])
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

func (detect Detector) evaluateStateFromTriggers(states map[string]interface{}) (string, bool) {
	log.Debugf("state: %+v", states)
	matched := false
	profiles := make(map[string]cfg.Profile)
	keys := make([]string, len(detect.profiles))
	i := 0
	for k := range detect.profiles {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		profiles[k] = detect.profiles[k]
	}
	for name, profile := range profiles {
		for key, value := range profile.Trigger {
			if _, ok := states[key]; ok {
				if fmt.Sprintf("%v", value) == fmt.Sprintf("%v", states[key]) {
					log.Debugf("Profile '%s' part '%s' matched.", name, key)
					matched = true
					continue
				}
				log.Debugf("Values: %+v == %+v\n", value, states[key])
				log.Debugf("Profile '%s' part '%s' NOT matched.", name, key)
				matched = false
				break
			}
		}
		if matched {
			log.Infof("Profile '%s' matched!", name)
			return name, true
		}
		log.Debugf("Profile '%s' didn't match with state.", name)
	}
	return "", false
}
