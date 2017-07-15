package config

type PollerConfig struct {
	PollInterval string `yaml:"pollInterval"`
	Timeout      string `yaml:"timeout"`
}
