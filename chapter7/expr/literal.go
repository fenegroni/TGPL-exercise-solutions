package expr

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
