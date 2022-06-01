package plot

import (
	"fmt"
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// F is the signature of a 3-D surface function.
type F func(x, y float64) float64

// Surface computes an SVG rendering of a 3-D surface function f.
// It returns the total number of bytes written and any write error encountered.
func Surface(w io.Writer, f F) (written int, err error) {
	var n int
	n, err = fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	written += n
	if err != nil {
		return
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			n, err = fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			written += n
		}
	}
	n, err = fmt.Fprintln(w, "</svg>")
	written += n
	return
}

func corner(i, j int, f F) (sx float64, sy float64) {
	// Find point (x,y) at Corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}
