package expr

import (
	"fmt"
	"math"
)

type min struct {
	x, y Expr
}

func (c min) Eval(env Env) float64 {
	return math.Min(c.x.Eval(env), c.y.Eval(env))
}

func (c min) Check(vars map[Var]bool) error {
	if err := c.x.Check(vars); err != nil {
		return err
	}
	return c.y.Check(vars)
}

func (c min) String() string {
	return fmt.Sprintf("min(%s, %s)", c.x, c.y)
}
