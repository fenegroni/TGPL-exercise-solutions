package plot

import (
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var

// TODO implement surface from TGPL pg59
func Surface(w io.Writer, f func(x, y float64) float64) {

}
