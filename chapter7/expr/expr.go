package expr

import "errors"

type Expr interface {
	Eval(env Env) float64
}

func Parse(_ string) (Expr, error) {
	return nil, errors.New("not implemented")
}
