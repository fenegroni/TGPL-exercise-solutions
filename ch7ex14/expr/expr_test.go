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
	for _, test := range tests {
		expr, err := Parse(test.exp)
		if err != nil {
			t.Errorf("Parse: %s", err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%q.Eval() in %v = %q, want %q",
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
		{"math.Pi", Env{}, fmt.Errorf("."), nil},
		{"!true", Env{}, fmt.Errorf("!"), nil},
		{"\"hello\"", Env{}, fmt.Errorf("\""), nil},
		{"log(10)", Env{}, nil, fmt.Errorf("log")},
		{"sqrt(1, 2)", Env{}, nil, fmt.Errorf("sqrt")},
	}
	for _, test := range tests {
		expr, err := Parse(test.exp)
		if test.wantParseErr == nil && err != nil {
			t.Errorf("%q.Parse() in %v returns unexpected error %v", test.exp, test.env, err)
			continue
		}
		if test.wantParseErr != nil {
			if err == nil {
				t.Errorf("%q.Parse() in %v must give error containing %q, got %v", test.exp, test.env, test.wantParseErr, err)
			}
			if !strings.Contains(err.Error(), test.wantParseErr.Error()) {
				t.Errorf("%q.Parse() in %v must give error containing %q, got %q", test.exp, test.env, test.wantParseErr, err)
			}
			continue
		}
		vars := make(map[Var]bool)
		err = expr.Check(vars)
		if err == nil {
			t.Errorf("%q.Check() must give error containing %q, got %v", test.exp, test.wantCheckErr, err)
			continue
		}
		if !strings.Contains(err.Error(), test.wantCheckErr.Error()) {
			t.Errorf("%q.Check() must give error containing %q, got %q", test.exp, test.wantCheckErr, err)
			continue
		}
	}
}
