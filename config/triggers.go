package config

type TriggersConfig struct {
	IPAddress *IPAddressConfig `omitempty,yaml:"ipaddress"`
	Xrandr    *XrandrConfig    `omitempty,yaml:"xrandr"`
}

type IPAddressConfig struct {
	Interfaces []string `yaml:"interfaces"`
}

type XrandrConfig struct {
	XrandrBinary    string   `yaml:"xrandrBinary"`
	ScreensToIgnore []string `yaml:"screensToIgnore"`
}
