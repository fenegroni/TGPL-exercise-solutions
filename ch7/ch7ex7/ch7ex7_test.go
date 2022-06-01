package ch7ex7

import (
	"flag"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7/ch7ex6/tempconv"
	"strings"
	"testing"
)

func TestCelsiusFlag_prettyPrint(t *testing.T) {
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	var out strings.Builder
	flags.SetOutput(&out)
	_ = tempconv.CelsiusFlag(flags, "temp", 20, "the temperature")
	flags.PrintDefaults()
	want := "(default 20°C)"
	got := out.String()
	if !strings.Contains(got, want) {
		t.Errorf("Defaults output does not contain %q, got %q", want, got)
	}
}
