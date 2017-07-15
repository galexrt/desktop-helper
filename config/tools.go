package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Read config file by filename and returns Config
func Read(filename string) (*Config, error) {
	config := &Config{
		PollerOptions: PollerOptions{
			PollInterval: 3,
			Timeout:      3,
		},
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}
	return config, nil
}