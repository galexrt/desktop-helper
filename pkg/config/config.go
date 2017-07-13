package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config contains the built config
type Config struct {
	Version         string             `yaml:"version"`
	DetectorOptions DetectorOptions    `yaml:"detectorOptions"`
	Profiles        map[string]Profile `yaml:"profiles"`
}

// DetectorOptions
type DetectorOptions struct {
	PollInterval int64 `yaml:"pollInterval"`
	Timeout      int64 `yaml:"timeout"`
}

// Profile contains enable and disable ProfileEvent options
type Profile struct {
	Name    string                            `yaml:"name"`
	Enable  map[string]map[string]interface{} `yaml:"on_enable"`
	Disable map[string]map[string]interface{} `yaml:"on_disable"`
	Trigger map[string]interface{}            `yaml:"trigger"`
}

// Read config file by filename and returns Config
func Read(filename string) (*Config, error) {
	config := &Config{
		DetectorOptions: DetectorOptions{
			PollInterval: 3,
			Timeout:      3,
		},
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Config{}, err
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("error: %v", err)
	}
	return config, nil
}
