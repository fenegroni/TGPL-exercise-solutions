package web

import (
	"fmt"
	"net/http"
)
import "github.com/fenegroni/TGPL-exercise-solutions/ch7ex13/expr"

func Plot(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	/* exp */ _, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	/* TODO implement surface
	surface(w, func(x, y float64) float64 {

	}*/
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
