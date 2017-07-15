package utils

import (
	"strings"
)

func GetTriggerName(name string) string {
	name = strings.Split(name, ".")[1]
	name = strings.Replace(name, "Option", "", 1)
	return strings.ToLower(name)
}
