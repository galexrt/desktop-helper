package triggers

import "context"

// Trigger contains function to match the given
type Trigger interface {
	GetState(context.Context, interface{}) (map[string]interface{}, error)
}

var triggers = make(map[string]Trigger)

// Register a trigger
func Register(name string, trigger Trigger) {
	triggers[name] = trigger
}

// Exist check if a trigger exists by name (key)
func Exist(name string) bool {
	_, ok := triggers[name]
	return ok
}

// Get
func Get(name string) Trigger {
	trigger, _ := triggers[name]
	return trigger
}
