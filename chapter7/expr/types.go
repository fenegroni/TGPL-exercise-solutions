package expr

type Var string

type Env map[Var]float64

type literal float64

type unary struct {
	op rune // '+', '-'
	x  Expr
}

type binary struct {
	op   rune // '+', '-', '*', '/'
	x, y Expr
}

type call struct {
	fn   string // "pow", "sin", "sqrt"
	args []Expr
}
