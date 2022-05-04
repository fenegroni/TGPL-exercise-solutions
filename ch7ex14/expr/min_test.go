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
	wantResult := literal(3).String()
	gotResult := literal(ex.Eval(nil)).String()
	if gotResult != wantResult {
		t.Errorf("Want %s, got %s", wantResult, gotResult)
	}
}
