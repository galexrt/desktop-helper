package exec

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/galexrt/desktop-helper/pkg/actions"
	"github.com/prometheus/common/log"
)

// Action contains options
type Action struct {
	actions.Action
}

func init() {
	actions.Register("exec", New())
}

// New create new ScreenLayout struct
func New() actions.Action {
	return &Action{}
}

// Run against the given options
func (action Action) Run(ctx context.Context, options map[string]interface{}) (string, error) {
	parts := strings.Fields(options["command"].(string))
	command := parts[0]
	parts = parts[1:len(parts)]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cmd *exec.Cmd
	if len(command) > 1 {
		cmd = exec.CommandContext(ctx, command, parts...)
	} else {
		cmd = exec.CommandContext(ctx, command, "")
	}
	out, err := cmd.CombinedOutput()
	log.With("action", "exec").With("cmd", options["command"]).With("err", err).Debugf("Output: %+s", string(out))
	return string(out), err
}
