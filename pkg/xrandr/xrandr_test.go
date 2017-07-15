package xrandr

import (
	"reflect"
	"testing"
)

var tests = []string{
	`Screen 0: minimum 8 x 8, current 1920 x 2160, maximum 32767 x 32767
eDP1 connected 1920x1080+0+1080 (normal left inverted right x axis y axis) 310mm x 170mm
   1920x1080     60.02*+  47.99
   1400x1050     59.98
   1600x900      60.00
   1280x1024     60.02
   1280x960      60.00
   1368x768      60.00
   1280x720      60.00
   1024x768      60.00
   1024x576      60.00
   960x540       60.00
   800x600       60.32    56.25
   864x486       60.00
   640x480       59.94
   720x405       60.00
   640x360       60.00
DP1 disconnected (normal left inverted right x axis y axis)
DP1-1 disconnected (normal left inverted right x axis y axis)
DP1-2 disconnected (normal left inverted right x axis y axis)
DP1-3 disconnected primary (normal left inverted right x axis y axis)
HDMI1 connected 1920x1080+0+0 (normal left inverted right x axis y axis) 530mm x 300mm
   1920x1080     60.00*+  59.94
   1680x1050     59.88
   1600x900      60.00
   1280x1024     75.02    60.02
   1280x800      59.91
   1152x864      75.00
   1280x720      60.00    59.94
   1024x768      75.03    60.00
   832x624       74.55
   800x600       75.00    60.32
   640x480       75.00    60.00    59.94
   720x400       70.08
HDMI2 disconnected (normal left inverted right x axis y axis)
VIRTUAL1 disconnected (normal left inverted right x axis y axis)`,
	`Screen 0: minimum 8 x 8, current 1920 x 2160, maximum 32767 x 32767
eDP1 disconnected (normal left inverted right x axis y axis)
DP1 disconnected (normal left inverted right x axis y axis)
DP1-1 disconnected (normal left inverted right x axis y axis)
DP1-2 disconnected (normal left inverted right x axis y axis)
DP1-3 disconnected primary (normal left inverted right x axis y axis)
HDMI1 disconnected (normal left inverted right x axis y axis)
HDMI2 disconnected (normal left inverted right x axis y axis)
VIRTUAL1 disconnected (normal left inverted right x axis y axis)`,
}
var results = []Screens{
	{
		ConnectedCount: 2,
		List: []string{
			"eDP1",
			"HDMI1",
		},
	},
	{
		ConnectedCount: 0,
		List:           []string{},
	},
}

func TestParse(t *testing.T) {
	for key, test := range tests {
		result, _ := Parse(test)
		if reflect.DeepEqual(result, results[key]) {
			t.Fatalf("result is different: %+v - %+v", result, results[key])
		}
	}
}
