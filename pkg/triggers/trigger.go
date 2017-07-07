package triggers

// Trigger contains function to match the given
type Trigger interface {
	Match(struct{}) (bool, error)
}

var triggers = make(map[string]Trigger)

// Register a trigger
func Register(name string, trigger Trigger) {
	triggers[name] = trigger
}
