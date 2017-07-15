package triggers

import (
	"fmt"
	"sync"

	"github.com/galexrt/desktop-helper/config"
)

type Ctor func(config.TriggersConfig) (Trigger, error)

type Trigger interface {
	GetState() error
	Match(config.TriggerOption) (bool, error)
}

type Manager struct {
	sync.Mutex
	config   config.TriggersConfig
	triggers map[string]Trigger
}

var constructors = make(map[string]Ctor)

func NewManager(cfg config.TriggersConfig) *Manager {
	return &Manager{
		config:   cfg,
		triggers: make(map[string]Trigger),
	}
}

func (mgr *Manager) Get(name string) (Trigger, error) {
	if trg, ok := mgr.triggers[name]; ok {
		return trg, nil
	}
	return mgr.newTrigger(name)
}

func (mgr *Manager) newTrigger(name string) (Trigger, error) {
	mgr.Mutex.Lock()
	ctor, ok := constructors[name]
	mgr.Mutex.Unlock()
	if !ok {
		return nil, fmt.Errorf("trigger with name '%s' not found.", name)
	}
	trg, err := ctor(mgr.config)
	if err != nil {
		return nil, err
	}
	mgr.triggers[name] = trg
	return trg, nil
}

func Register(name string, ctor Ctor) {
	constructors[name] = ctor
}
