package actions

// Action contains function to match the given
type Action interface {
	Run(interface{}) error
}

var actions = make(map[string]Action)

// Register an action
func Register(name string, action Action) {
	actions[name] = action
}
