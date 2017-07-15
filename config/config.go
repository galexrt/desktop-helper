package config

// Config contains the config file structure
type Config struct {
	Version        string         `yaml:"version"`
	PollerConfig   PollerConfig   `yaml:"pollerconfig"`
	TriggersConfig TriggersConfig `yaml:"triggersconfig"`
	ActionsConfig  ActionsConfig  `yaml:"actionsconfig"`
	RunnerConfig   RunnerConfig   `yaml:"runnerconfig"`
	Profiles       []Profile      `yaml:"profiles"`
}
