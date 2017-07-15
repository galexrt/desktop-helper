package config

type ActionsConfig struct {
	Exec      ExecConfig      `yaml:"exec"`
	Libnotify LibnotifyConfig `yaml:"libnotify"`
}

type ExecConfig struct {
}

type LibnotifyConfig struct {
}
