package ch7ex13

import (
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"math"
	"testing"
)

func TestStringEval(t *testing.T) {
	tests := []struct {
		exp string
		env expr.Env
	}{
		{"sqrt(A / pi)", expr.Env{"A": 87616, "pi": math.Pi}},
		{"pow(x, 3) + pow(y, 3)", expr.Env{"x": 12, "y": 1}},
		{"pow(x, 3) + pow(y, 3)", expr.Env{"x": 9, "y": 10}},
		{"5 / 9 * (F - 32)", expr.Env{"F": -40}},
		{"5 / 9 * (F - 32)", expr.Env{"F": 32}},
		{"5 / 9 * (F - 32)", expr.Env{"F": 212}},
	}
	for _, test := range tests {
		exp, err := expr.Parse(test.exp)
		if err != nil {
			t.Errorf("Parse: %s", err)
			continue
		}
		result1 := fmt.Sprintf("%.6g", exp.Eval(test.env))
		got := exp.String()
		exp2, err := expr.Parse(got)
		if err != nil {
			t.Errorf("Parse: %s", err)
			continue
		}
		result2 := fmt.Sprintf("%.6g", exp2.Eval(test.env))
		t.Logf("%s => %s", test.exp, got)
		if result1 != result2 {
			t.Errorf("%q and %q don't give the same result", test.exp, got)
		}
	}
}

// TODO implement test option 2

// TODO implement test option 3
