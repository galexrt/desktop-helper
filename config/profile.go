package config

type Profile struct {
	Name    string        `yaml:"name"`
	Enable  ActionOption  `yaml:"onEnable"`
	Disable ActionOption  `yaml:"onDisable"`
	Trigger TriggerOption `yaml:"trigger"`
}

type ActionOption struct {
	Exec      ExecOption      `yaml:"exec"`
	Libnotify LibnotifyOption `yaml:"libnotify"`
}

type ExecOption struct {
	Command string `yaml:"command"`
}

type LibnotifyOption struct {
	Urgency int    `yaml:"urgency"`
	Delay   string `yaml:"delay"`
	Title   string `yaml:"title"`
	Message string `yaml:"message"`
	Image   string `yaml:"image"`
}

// TRIGGER OPTIONS ===============

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
