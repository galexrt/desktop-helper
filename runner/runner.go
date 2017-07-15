package runner

import (
	"context"
	"reflect"
	"sync"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/runner/actions"
	"github.com/galexrt/desktop-helper/utils"
)

type Runner struct {
	config     config.RunnerConfig
	actionsMgr *actions.Manager
}

func New(cfg config.RunnerConfig, actionsMgr *actions.Manager) (*Runner, error) {
	return &Runner{
		config:     cfg,
		actionsMgr: actionsMgr,
	}, nil
}

func (rner *Runner) OnEnable(ctx context.Context, profile config.Profile) error {
	return rner.runActions(ctx, profile.Enable)
}

func (rner *Runner) OnDisable(ctx context.Context, profile config.Profile) error {
	return rner.runActions(ctx, profile.Disable)
}

func (rner *Runner) runActions(ctx context.Context, actns config.ActionOption) error {
	errors := make(chan error, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		msValuePtr := reflect.ValueOf(&actns)
		msValue := msValuePtr.Elem()
		for i := 0; i < msValue.NumField(); i++ {
			field := msValue.Field(i)
			name := utils.GetTriggerName(field.Type().String())
			actn, err := rner.actionsMgr.Get(name)
			if err != nil {
				errors <- err
				return
			}
			wg.Add(1)
			go func(action actions.Action) {
				defer wg.Done()
				err = action.Execute(actns)
				if err != nil {
					errors <- err
				}
			}(actn)
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
