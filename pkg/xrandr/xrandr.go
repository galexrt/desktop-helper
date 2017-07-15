package xrandr

import "strings"

type Screens struct {
	ConnectedCount int
	List           []string
}

func Parse(input string) (*Screens, error) {
	screens := &Screens{
		ConnectedCount: 0,
	}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "connected") {
			screens.List = append(screens.List, strings.SplitN(line, " ", 1)[0])
			screens.ConnectedCount++
		}
	}
	return screens, nil
}
