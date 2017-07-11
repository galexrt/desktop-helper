package exec

import (
	"context"
	"os/exec"
	"strings"

	"github.com/galexrt/desktop-helper/pkg/actions"
	"github.com/prometheus/common/log"
)

// Action contains options
type Action struct {
	actions.Action
}

// ActionOptions
type ActionOptions struct {
	Command string
}

func init() {
	actions.Register("exec", New())
}

// New create new ScreenLayout struct
func New() actions.Action {
	return &Action{}
}

// Run against the given options
func (action Action) Run(opts interface{}) error {
	options := opts.(ActionOptions)
	command := []string{"echo", "Hello World!"}
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	cmd := exec.CommandContext(ctx, command[0], strings.Join(command[1:], " "))
	err := cmd.Run()
	if err != nil {
		return nil
	}
	out, err := cmd.CombinedOutput()
	log.With("action", "exec").With("command", options.Command).Debugf("Output: %+s\n", string(out))
	return err
}
