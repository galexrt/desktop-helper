package exec

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/runner/actions"
	log "github.com/sirupsen/logrus"
)

type Action struct {
	actions.Action
	cfg *config.ExecConfig
}

func init() {
	actions.Register("exec", New)
}

func New(cfg config.ActionsConfig) (actions.Action, error) {
	return &Action{
		cfg: cfg.Exec,
	}, nil
}

func (exe *Action) Execute(opts config.ActionOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cmd *exec.Cmd
	if len(opts.Exec.Args) > 0 {
		cmd = exec.CommandContext(ctx, opts.Exec.Command, opts.Exec.Args...)
	} else {
		cmd = exec.CommandContext(ctx, opts.Exec.Command, "")
	}
	out, err := cmd.CombinedOutput()
	log.WithField("action", "exec").WithField("cmd", opts.Exec.Command).
		WithField("err", err).Debugf("Output: %+s", prepOutput(string(out)))
	return err
}

func prepOutput(out string) string {
	out = strings.Replace(out, "\n", "\\n", -1)
	out = strings.Replace(out, "\t", "\\t", -1)
	out = strings.Replace(out, "\r", "\\r", -1)
	return out
}
