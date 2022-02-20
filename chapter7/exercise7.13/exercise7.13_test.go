package exercise7_13

import "testing"

func TestExpr(t *testing.T) {
	e1 := NewExpr("sqrt(A / pi)")
	e2 := NewExpr("pow(x, 3) + pow(y, 3)")
	e3 := NewExpr("(F - 32) * 5 / 9")
	t.Logf("e1: %v", e1)
	t.Logf("e2: %v", e2)
	t.Logf("e3: %v", e3)
}
