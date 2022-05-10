package ch7ex15

import (
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"testing"
)

func TestVarsDetected(t *testing.T) {
	tests := []string{
		"x / 2",
	}
	for _, test := range tests {
		ex, err := expr.Parse(test)
		if err != nil {
			t.Errorf("Parse(%q): %s", test, err)
			continue
		}
		vars := make(map[expr.Var]bool)
		err = ex.Check(vars)
		if err != nil {
			t.Errorf("Check(%q): %s", test, err)
			continue
		}
	}
}

func ExampleCh7ex15() {
	// TODO validate an Example can read stdin, write stdout
}
