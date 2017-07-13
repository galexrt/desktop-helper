package runner

import (
	"context"
	"sync"

	"github.com/galexrt/desktop-helper/pkg/actions"
	"github.com/galexrt/desktop-helper/pkg/config"
	"github.com/prometheus/common/log"
)

// Runner
type Runner struct {
	profiles     map[string]config.Profile
	lastProfile  string
	actionsCache map[string]actions.Action
}

func New(config *config.Config) *Runner {
	return &Runner{
		profiles:     config.Profiles,
		actionsCache: make(map[string]actions.Action),
	}
}

// OnEnable
func (run *Runner) OnEnable(ctx context.Context, profile string) error {
	if run.lastProfile == profile {
		return nil
	}
	run.lastProfile = profile
	log.Debugf("Enable for profile '%s' triggered.", profile)
	_, err := run.runActions(ctx, run.profiles[run.lastProfile].Enable)
	return err

}

// OnDisable
func (run *Runner) OnDisable(ctx context.Context, profile string) error {
	log.Debugf("Disable for profile '%s' triggered.", profile)
	_, err := run.runActions(ctx, run.profiles[run.lastProfile].Disable)
	return err
}

func (run *Runner) runActions(ctx context.Context, list map[string]map[string]interface{}) (map[string]interface{}, error) {
	wg := sync.WaitGroup{}
	tctx, cancel := context.WithCancel(ctx)
	errors := make(chan error)
	outputs := map[string]interface{}{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var mutex = &sync.Mutex{}
		for name, conf := range list {
			if _, ok := run.actionsCache[name]; !ok {
				run.actionsCache[name] = actions.Get(name)
			}
			wg.Add(1)
			go func(name string, action actions.Action, actionConf map[string]interface{}) {
				defer wg.Done()
				log.Debugf("action: %+v; actionConf: %+v", name, actionConf)
				output, err := action.Run(tctx, actionConf)
				mutex.Lock()
				outputs[name] = output
				mutex.Unlock()
				if err != nil {
					errors <- err
				}
			}(name, run.actionsCache[name], conf)
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

	return outputs, err
}
