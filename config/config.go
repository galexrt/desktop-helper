package config

// Config contains the config file structure
type Config struct {
	Version        string         `yaml:"version"`
	PollerConfig   PollerConfig   `yaml:"pollerConfig"`
	TriggersConfig TriggersConfig `yaml:"triggersConfig"`
	ActionsConfig  ActionsConfig  `yaml:"actionsConfig"`
	RunnerConfig   RunnerConfig   `yaml:"runnerConfig"`
	Profiles       []Profile      `yaml:"profiles"`
}
