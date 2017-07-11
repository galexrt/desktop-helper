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

// Exist check if an action exists by name (key)
func Exist(name string) bool {
	_, ok := actions[name]
	return ok
}
