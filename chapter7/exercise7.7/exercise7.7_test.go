package exercise7_7

import (
	"TGPL-exercise-solutions/chapter7/exercise7.6/tempconv"
	"flag"
	"strings"
	"testing"
)

func TestCelsiusFlag_prettyPrint(t *testing.T) {
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	var out strings.Builder
	flags.SetOutput(&out)
	_ = tempconv.CelsiusFlag(flags, "temp", 0, "the temperature")
	flags.PrintDefaults()
	want := "(default 0°C)"
	got := out.String()
	if !strings.Contains(got, want) {
		t.Errorf("Defaults output does not contain %q, got %q", want, got)
	}
}
