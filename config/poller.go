package config

type PollerOptions struct {
	PollInterval int64 `yaml:"pollInterval"`
	Timeout      int64 `yaml:"timeout"`
}
