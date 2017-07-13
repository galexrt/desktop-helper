package actions

import "context"

// Action contains function to match the given
type Action interface {
	Run(context.Context, map[string]interface{}) (string, error)
}

var actions = make(map[string]Action)

// Register an action
func Register(name string, action Action) {
	actions[name] = action
}

// Exist check if an action exists by name (key)
func Exist(name string) bool {
	_, ok := actions[name]
	return ok
}

// Get
func Get(name string) Action {
	action, _ := actions[name]
	return action
}
