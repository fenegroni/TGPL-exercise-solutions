package expr

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		exp  string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.exp != prevExpr {
			fmt.Printf("\n%s\n", test.exp)
			prevExpr = test.exp
		}
		expr, err := Parse(test.exp)
		if err != nil {
			t.Errorf("Parse: %s", err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q",
				test.exp, test.env, got, test.want)
		}
	}
}

func TestParseAndCheckErrors(t *testing.T) {
	tests := []struct {
		exp          string
		env          Env
		wantParseErr error
		wantCheckErr error
	}{
		{"x % 2", Env{"x": 42}, fmt.Errorf("%%"), nil},
		{"x + 2", Env{"y": 42}, nil, fmt.Errorf("x")},
	}
	var prevExpr string
	for _, test := range tests {
		if test.exp != prevExpr {
			fmt.Printf("\n%s\n", test.exp)
			prevExpr = test.exp
		}
		expr, err := Parse(test.exp)
		if test.wantParseErr == nil && err != nil {
			t.Errorf("%s.Parse() in %v returns unexpected error %q", test.exp, test.env, err)
			continue
		}
		if test.wantParseErr != nil {
			if err == nil {
				t.Errorf("%s.Parse() in %v does not error, want error containing %q", test.exp, test.env, test.wantParseErr)
			}
			if !strings.Contains(err.Error(), test.wantParseErr.Error()) {
				t.Errorf("%s.Parse() in %v = error %q, can't find %q", test.exp, test.env, err, test.wantParseErr)
			}
			continue
		}
		vars := make(map[Var]bool)
		err = expr.Check(vars)
		if err == nil {
			t.Errorf("%s.Check() in %v does not error, want error containing %q", test.exp, test.env, test.wantCheckErr)
			continue
		}
		if !strings.Contains(err.Error(), test.wantParseErr.Error()) {
			t.Errorf("%s.Check() in %v = error %q, can't find %q", test.exp, test.env, err, test.wantCheckErr)
			continue
		}
		// TODO check vars against env? maybe need to have nil errors for Check too
	}
}
