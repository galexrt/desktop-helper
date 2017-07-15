package libnotify

import (
	"fmt"
	"os"
	"time"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/runner/actions"

	notify "github.com/mqu/go-notify"
)

type Action struct {
	actions.Action
	cfg *config.LibnotifyConfig
}

func init() {
	actions.Register("libnotify", New)
}

func New(cfg config.ActionsConfig) (actions.Action, error) {
	return &Action{
		cfg: cfg.Libnotify,
	}, nil
}

func (exe *Action) Execute(opts config.ActionOption) error {
	duration, err := time.ParseDuration(opts.Libnotify.Delay)
	if err != nil {
		return err
	}
	if opts.Libnotify.Urgency > 2 || opts.Libnotify.Urgency < 0 {
		return fmt.Errorf("invalid libnotify urgency: '%d'", opts.Libnotify.Urgency)
	}

	notify.Init(opts.Libnotify.Message)
	defer notify.UnInit()

	image := ""
	if opts.Libnotify.Image != "" {
		image = opts.Libnotify.Message
	}
	hello := notify.NotificationNew(opts.Libnotify.Title,
		opts.Libnotify.Message,
		image)
	hello.SetUrgency(notify.NotifyUrgency(opts.Libnotify.Urgency))

	if hello == nil {
		return fmt.Errorf("Unable to create a new notification: %+v", hello)
	}

	if e := notify.NotificationShow(hello); e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e.Message())
	}

	time.Sleep(duration)
	notify.NotificationClose(hello)

	return nil
}
