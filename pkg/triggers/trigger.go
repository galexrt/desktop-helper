package triggers

import "context"

// Trigger contains function to match the given
type Trigger interface {
	GetState(context.Context, interface{}) (bool, error)
}

var triggers = make(map[string]Trigger)

// Register a trigger
func Register(name string, trigger Trigger) {
	triggers[name] = trigger
}

func Exist(name string) bool {
	_, ok := triggers[name]
	return ok
}

func Get(name string) Trigger {
	trigger, _ := triggers[name]
	return trigger
}
