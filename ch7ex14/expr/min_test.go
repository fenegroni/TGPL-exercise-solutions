package expr

import "testing"

func TestMinWithSyntaxTree(t *testing.T) {
	asTree := min{literal(2), literal(1)}
	wantResult := literal(1).String()
	gotResult := literal(asTree.Eval(nil)).String()
	if gotResult != wantResult {
		t.Errorf("Wrong result: want %s, got %s", wantResult, gotResult)
	}
}
