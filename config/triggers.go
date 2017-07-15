package config

type TriggersConfig struct {
	IPAddress *IPAddressConfig `omitempty,yaml:"ipaddress"`
	Xrandr    *XrandrConfig    `omitempty,yaml:"xrandr"`
}

type IPAddressConfig struct {
	Interfaces []string `yaml:"interfaces"`
}

type XrandrConfig struct {
	XrandrBinary   string   `yaml:"xrandrbinary"`
	XAuthoritiy    string   `yaml:"xauthoritiy"`
	IgnoreSegFault bool     `yaml:"ignoresegfault"`
	IgnoreErrors   bool     `yaml:"ignoreerrors"`
	Display        string   `yaml:"display"`
	ScreensIgnore  []string `yaml:"screensignore"`
}
