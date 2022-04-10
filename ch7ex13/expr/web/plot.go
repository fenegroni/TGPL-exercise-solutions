package web

import (
	"fmt"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/plot"
	"math"
	"net/http"
)

func Plot(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	exp, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	_, _ = plot.Surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return exp.Eval(expr.Env{"x": x, "y": y, "r": r})
	})
}

func parseAndCheck(s string) (expr.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	exp, err := expr.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[expr.Var]bool)
	if err := exp.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return exp, nil
}
