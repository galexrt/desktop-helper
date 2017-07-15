package poller

import (
	"context"
	"reflect"
	"sync"
	"time"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/poller/triggers"
	"github.com/galexrt/desktop-helper/runner"
	"github.com/galexrt/desktop-helper/runner/actions"
	"github.com/galexrt/desktop-helper/utils"
	log "github.com/sirupsen/logrus"
)

type Poller struct {
	triggersMgr     *triggers.Manager
	runner          *runner.Runner
	profiles        []config.Profile
	profile         int
	enabledTriggers map[string]struct{}
}

func New(cfg *config.Config) (*Poller, error) {
	actionsMgr := actions.NewManager(cfg.ActionsConfig)
	rner, err := runner.New(cfg.RunnerConfig, actionsMgr)
	if err != nil {
		return nil, err
	}
	poller := &Poller{
		triggersMgr:     triggers.NewManager(cfg.TriggersConfig),
		runner:          rner,
		profiles:        cfg.Profiles,
		profile:         -1,
		enabledTriggers: getEnabledTriggers(cfg.Profiles),
	}
	return poller, nil
}

func (pol *Poller) Run(bctx context.Context) error {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(bctx)
	errors := make(chan error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Debugf("match triggers")
			// CHECK STATE AND MATCH
			profile, err := pol.matchTriggers(ctx)
			if err != nil {
				errors <- err
				return
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
			// CALL ACTIONS
			if pol.profile != profile {
				log.WithField("id", profile).Debugf("new profile match")
				if pol.profile != -1 {
					err = pol.runner.OnDisable(ctx, pol.profiles[pol.profile])
					if err != nil {
						errors <- err
						return
					}
				}
				if profile != -1 {
					err = pol.runner.OnEnable(ctx, pol.profiles[profile])
					if err != nil {
						errors <- err
						return
					}
				}
				pol.profile = profile
			} else {
				log.WithField("id", profile).Debugf("no profile match change")
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(3 * time.Second):
			}
		}
	}()
	wgc := make(chan struct{})
	go func() {
		defer close(wgc)
		wg.Wait()
	}()
	var err error
	select {
	case <-wgc:
	case <-bctx.Done():
	case <-ctx.Done():
	case err = <-errors:
	}
	cancel()
	wg.Wait()
	return err
}

func (pol *Poller) getState(ctx context.Context) error {
	errors := make(chan error, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for name := range pol.enabledTriggers {
			trigger, err := pol.triggersMgr.Get(name)
			if err != nil {
				errors <- err
			}
			wg.Add(1)
			go func(trg triggers.Trigger) {
				defer wg.Done()
				err = trg.GetState()
				if err != nil {
					errors <- err
				}
			}(trigger)
		}
	}()
	wgc := make(chan struct{})
	go func() {
		defer close(wgc)
		wg.Wait()
	}()
	var err error
	select {
	case <-wgc:
	case <-ctx.Done():
	case err = <-errors:
	}
	wg.Wait()
	return err
}

func (pol *Poller) matchTriggers(ctx context.Context) (int, error) {
	profID := -1
	err := pol.getState(ctx)
	if err != nil {
		return profID, err
	}
	var match bool
	for id, profile := range pol.profiles {
		msValuePtr := reflect.ValueOf(&profile.Trigger)
		msValue := msValuePtr.Elem()

		for i := 0; i < msValue.NumField(); i++ {
			field := msValue.Field(i)
			if !field.IsNil() {
				name := utils.GetTriggerName(field.Type().String())
				trigger, err := pol.triggersMgr.Get(name)
				if err != nil {
					return profID, err
				}
				match, err = trigger.Match(profile.Trigger)
				if err != nil {
					return profID, err
				}
				log.WithField("trigger", name).Debugf("match: %+v", match)
			}
		}
		if match {
			profID = id
			break
		}
	}
	if !match {
		profID = -1
	}
	return profID, nil
}

func getEnabledTriggers(profiles []config.Profile) map[string]struct{} {
	enabledTriggers := map[string]struct{}{}
	for _, profile := range profiles {
		msValuePtr := reflect.ValueOf(&profile.Trigger)
		msValue := msValuePtr.Elem()

		for i := 0; i < msValue.NumField(); i++ {
			field := msValue.Field(i)
			if !field.IsNil() {
				enabledTriggers[utils.GetTriggerName(field.Type().String())] = struct{}{}
			}
		}
	}
	return enabledTriggers
}
