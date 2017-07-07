package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config contains the built config
type Config struct {
	Version  string             `json:"version"`
	Profiles []Profile          `json:"profiles"`
	Triggers map[string]Trigger `json:"triggers"`
}

// Profile contains enable and disable ProfileEvent options
type Profile struct {
	Name    string       `json:"name"`
	Enable  ProfileEvent `json:"on_enable"`
	Disable ProfileEvent `json:"on_disable"`
}

// ProfileEvent contains on enable/disable options
type ProfileEvent struct {
	Screenlayout string `json:"screenlayout"`
	Exec         string `json:"exec"`
}

// Trigger contains a trigger configuration
type Trigger struct {
	Network []string `json:"network"`
	Screens Screens  `json:"screens"`
}

// Screens contains screens trigger config
type Screens struct {
	Count     int      `json:"count"`
	Connected []string `json:"connected"`
}

// Read config file by filename and returns Config
func Read(filename string) (*Config, error) {
	config := &Config{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Config{}, nil
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("error: %v", err)
	}
	return config, nil
}
