package expr

type Expr interface {
	Eval(env Env) float64
}
