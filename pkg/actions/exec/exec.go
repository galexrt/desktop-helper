package exec

import (
	"io/ioutil"
	"os/exec"
	"strings"

	"golang.org/x/net/context"

	"github.com/galexrt/desktop-helper/pkg/actions"
	"github.com/prometheus/common/log"
)

// ExecAction contains options
type ExecAction struct {
	actions.Action
}

type ExecActionOptions struct {
	Command string
}

func init() {
	actions.Register("exec", NewExecAction())
}

// NewScreenLayout create new ScreenLayout struct
func NewExecAction() actions.Action {
	return &ExecAction{}
}

// Run against the given options
func (execAction ExecAction) Run(config interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 10)
	command := []string{"echo", "Hello World!"}
	cmd := exec.CommandContext(ctx, command[0], strings.Join(command[1:], " "))
	err := cmd.Run()
	stdout, _ := ioutil.ReadAll(cmd.Stdout)
	stderr, _ := ioutil.ReadAll(cmd.Stderr)
	log.With("action", "exec").With("command", options.Command).Debugf("stdout", string(stdout))
	log.With("action", "exec").With("command", options.Command).Debugf("stderr", string(stderr))
	return err
}
