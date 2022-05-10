package ch7ex15

import (
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"reflect"
	"testing"
)

func TestVarsDetected(t *testing.T) {
	type varMap map[expr.Var]bool
	tests := []struct {
		ex   string
		want varMap
	}{
		{"sin(a)", varMap{"a": true}},
		{"x / Y", varMap{"Y": true, "x": true}},
		{"X / y", varMap{"X": true, "y": true}},
		{"1+1", varMap{}},
		{"A + B * C / D * (A + B)", varMap{"A": true, "B": true, "C": true, "D": true}},
	}
	for _, test := range tests {
		ex, err := expr.Parse(test.ex)
		if err != nil {
			t.Errorf("Parse(%q): %s", test.ex, err)
			continue
		}
		got := make(varMap)
		err = ex.Check(got)
		if err != nil {
			t.Errorf("Check(%q): %s", test.ex, err)
			continue
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("variables in %q: want %v, got %v", test.ex, test.want, got)
		}
	}
}

func ExampleCh7ex15() {
	var x string
	fmt.Scanf("%s", &x)
	fmt.Printf("%s", x)
}
