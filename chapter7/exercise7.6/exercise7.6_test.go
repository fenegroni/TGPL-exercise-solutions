package exercise7_6

import (
	tempconv "TGPL-exercise-solutions/chapter7/exercise7.6/tempconv"
	"flag"
	"testing"
)

func TestCelsiusFlag(t *testing.T) {
	tests := []struct {
		line  []string
		value tempconv.Celsius
	}{
		{[]string{}, 20.0},
		{[]string{"-temp", "10C"}, 10.0},
		{[]string{"-temp", "10°C"}, 10.0},
		{[]string{"-temp", "-10°C"}, -10.0},
		{[]string{"-temp", "32F"}, 0},
		{[]string{"-temp", "32°F"}, 0},
	}
	for _, test := range tests {
		flags := flag.NewFlagSet("test", flag.ContinueOnError)
		var temp = tempconv.CelsiusFlag(flags, "temp", 20.0, "the temperature")
		if err := flags.Parse(test.line); err != nil {
			t.Error(err)
		}
		if *temp != test.value {
			t.Errorf("%v = %f, want %f", test.line, *temp, test.value)
		}
	}
}
