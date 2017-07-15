package actions

import (
	"fmt"
	"sync"

	"github.com/galexrt/desktop-helper/config"
)

type Ctor func(config.ActionsConfig) (Action, error)

type Action interface {
	Execute(config.ActionOption) error
}

type Manager struct {
	sync.Mutex
	config  config.ActionsConfig
	actions map[string]Action
}

var constructors = make(map[string]Ctor)

func NewManager(cfg config.ActionsConfig) *Manager {
	return &Manager{
		config:  cfg,
		actions: make(map[string]Action),
	}
}

func (mgr *Manager) Get(name string) (Action, error) {
	if trg, ok := mgr.actions[name]; ok {
		return trg, nil
	}
	return mgr.newAction(name)
}

func (mgr *Manager) newAction(name string) (Action, error) {
	mgr.Mutex.Lock()
	ctor, ok := constructors[name]
	mgr.Mutex.Unlock()
	if !ok {
		return nil, fmt.Errorf("action with name '%s' not found.", name)
	}
	trg, err := ctor(mgr.config)
	if err != nil {
		return nil, err
	}
	mgr.actions[name] = trg
	return trg, nil
}

func Register(name string, ctor Ctor) {
	constructors[name] = ctor
}
