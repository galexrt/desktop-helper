package xrandr

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/galexrt/desktop-helper/config"
	xrandrlib "github.com/galexrt/desktop-helper/pkg/xrandr"
	"github.com/galexrt/desktop-helper/poller/triggers"
)

type Trigger struct {
	triggers.Trigger
	cfg   *config.XrandrConfig
	state *xrandrlib.Screens
}

type IPAddress struct {
	Count int `yaml:"count"`
}

func init() {
	triggers.Register("xrandr", New)
}

func New(cfg config.TriggersConfig) (triggers.Trigger, error) {
	return &Trigger{
		cfg: cfg.Xrandr,
	}, nil
}

func (trg *Trigger) GetState() error {
	// TODO
	cmd := exec.Command(trg.cfg.XrandrBinary, "--query")
	env := os.Environ()
	env = append(env,
		fmt.Sprintf("DISPLAY=%s", trg.cfg.Display),
		fmt.Sprintf("XAUTHORITY=%s", trg.cfg.XAuthoritiy),
	)
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	screens, err := xrandrlib.Parse(string(out))
	trg.state = screens
	return err
}

func (trg *Trigger) Match(opts config.TriggerOption) (bool, error) {
	var match bool
	if trg.state.ConnectedCount == opts.Xrandr.ConnectedCount {
		match = true
	}
	if len(trg.state.List) == len(opts.Xrandr.Screens) {
		screens := make(map[string]struct{})
		for _, v := range trg.state.List {
			screens[v] = struct{}{}
		}
		for _, screen := range opts.Xrandr.Screens {
			if _, ok := screens[screen]; ok {
				match = true
			} else {
				return false, nil
			}
		}
	}
	return match, nil
}
