package expr

type Var string

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}
