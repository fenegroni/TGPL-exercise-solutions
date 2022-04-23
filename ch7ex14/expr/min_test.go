package expr

import "testing"

func TestMinWithSyntaxTree(t *testing.T) {
	asTree := min{2, 1}
	wantResult := 1
	gotResult := asTree.Eval(nil)
	if gotResult != wantResult {
		t.Errorf("Wrong result: want %d, got %d", wantResult, gotResult)
	}
}
