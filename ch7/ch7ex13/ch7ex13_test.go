package ch7ex13

import (
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7/ch7ex13/expr"
	"math"
	"reflect"
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

func TestStringParse(t *testing.T) {
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
	}
	for _, test := range tests {
		var exp expr.Expr
		var err error
		exp, err = expr.Parse(test)
		if err != nil {
			t.Errorf("Parse original exp: %s", err)
			continue
		}
		got1 := exp.String()
		exp, err = expr.Parse(got1)
		if err != nil {
			t.Errorf("Parse first string: %s", err)
			continue
		}
		got2 := exp.String()
		if got1 != got2 {
			t.Errorf("got1 %q != got2: %q", got1, got2)
		}
	}
}

func TestStringDeepCompare(t *testing.T) {
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
	}
	for _, test := range tests {
		var exp1, exp2 expr.Expr
		var err error
		exp1, err = expr.Parse(test)
		if err != nil {
			t.Errorf("Parse original exp: %s", err)
			continue
		}
		exp2, err = expr.Parse(exp1.String())
		if err != nil {
			t.Errorf("Parse string rep.: %s", err)
			continue
		}
		if false == reflect.DeepEqual(exp1, exp2) {
			t.Error("syntax trees are different")
		}
	}
}
