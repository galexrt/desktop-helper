package config

type TriggersConfig struct {
	IPAddress IPAddressConfig `yaml:"ipAddress"`
}

type IPAddressConfig struct {
	Interfaces []string `yaml:"interfaces"`
}
