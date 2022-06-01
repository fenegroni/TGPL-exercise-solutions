package expr

import "testing"

func TestMinSimpleSyntaxTree(t *testing.T) {
	asTree := min{literal(2), literal(1)}
	wantResult := literal(1).String()
	gotResult := literal(asTree.Eval(nil)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}

func TestMinSyntaxTreeUsingVar(t *testing.T) {
	env := Env{"A": 2, "B": 1}
	asTree := min{Var("A"), Var("B")}
	wantResult := literal(1).String()
	gotResult := literal(asTree.Eval(env)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}

func TestMinComplexTree(t *testing.T) {
	env := Env{"A": 2, "B": 3}
	asTree :=
		min{
			binary{'+',
				min{Var("A"), Var("B")},
				literal(1),
			},
			binary{'+',
				min{Var("A"), Var("B")},
				literal(2),
			},
		}
	wantResult := literal(3).String()
	gotResult := literal(asTree.Eval(env)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}

func TestParseMin(t *testing.T) {
	ex, err := Parse("min(2, 3)")
	if err != nil {
		t.Fatalf("Parse error: %s", err)
	}
	if err = ex.Check(nil); err != nil {
		t.Fatalf("Check error: %s", err)
	}
	wantResult := literal(2).String()
	gotResult := literal(ex.Eval(nil)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}

func TestParseMinComplexExpression(t *testing.T) {
	env := Env{"A": 2, "B": 3}
	ex, err := Parse("min(min(A, B) + 1, min(A, B) + 2)")
	if err != nil {
		t.Fatalf("Parse error: %s", err)
	}
	if err = ex.Check(map[Var]bool{"A": false, "B": false}); err != nil {
		t.Fatalf("Check error: %s", err)
	}
	wantResult := literal(3).String()
	gotResult := literal(ex.Eval(env)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}

func TestParseMinTooManyArgs(t *testing.T) {
	_, err := Parse("min(2, 3, 4)")
	t.Logf("Parse error: %v", err)
	if err == nil {
		t.Error("Parse did not detect error")
	}
}

func TestParseMinTooFewArgs(t *testing.T) {
	_, err := Parse("min(4)")
	t.Logf("Parse error: %v", err)
	if err == nil {
		t.Error("Parse did not detect error")
	}
	_, err = Parse("min()")
	t.Logf("Parse error: %v", err)
	if err == nil {
		t.Error("Parse did not detect error")
	}
}
