package tempconv

import (
	"flag"
	"fmt"
)

// *celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(set *flag.FlagSet, name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	set.Var(&f, name, usage)
	return &f.Celsius
}
