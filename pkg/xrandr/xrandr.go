package xrandr

import (
	"fmt"
	"strings"
)

type Screens struct {
	ConnectedCount int
	List           []string
}

func Parse(input string) (Screens, error) {
	screens := &Screens{
		ConnectedCount: 0,
	}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, " connected ") {
			screens.List = append(screens.List, strings.Split(line, " ")[0])
			screens.ConnectedCount++
		}
	}
	return *screens, nil
}

func ParseActiveMonitors(input string) (Screens, error) {
	screens := &Screens{
		ConnectedCount: 0,
	}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, ": +") {
			disp := strings.Split(line, "+")[1]
			screens.List = append(screens.List, strings.Split(disp, " ")[0])
			screens.ConnectedCount++
		}
	}
	fmt.Printf("TEST: %+v\n", screens)
	return *screens, nil
}
