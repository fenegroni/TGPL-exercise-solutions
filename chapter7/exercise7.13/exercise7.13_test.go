package exercise7_13

import (
	"github.com/fenegroni/TGPL-exercise-solutions/chapter7/expr"
	"testing"
)

func TestExpr(t *testing.T) {
	e1, _ := expr.Parse("sqrt(A / pi)")
	e2, _ := expr.Parse("pow(x, 3) + pow(y, 3)")
	e3, _ := expr.Parse("(F - 32) * 5 / 9")
	t.Logf("e1: %v", e1)
	t.Logf("e2: %v", e2)
	t.Logf("e3: %v", e3)
}
