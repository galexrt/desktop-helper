package config

type Profile struct {
	Name    string        `yaml:"name"`
	Enable  ActionOption  `yaml:"onEnable"`
	Disable ActionOption  `yaml:"onDisable"`
	Trigger TriggerOption `yaml:"trigger"`
}

type ActionOption struct {
	Exec ExecOption `yaml:"exec"`
}

type ExecOption struct {
	Command string `yaml:"command"`
}

type TriggerOption struct {
	IPAddress IPAddressOption `yaml:"ipAddress"`
}

type IPAddressOption struct {
	Addresses map[string]IPAddress `yaml:"addresses"`
}

type IPAddress struct {
	Address string `yaml:"address"`
	Key     int    `yaml:"key"`
}
