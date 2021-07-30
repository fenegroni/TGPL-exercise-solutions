package geometry

import "math"

type Point struct{ X, Y float64 }

// Distance between Points p and q
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy returnes the point coordinates scaled by factor
func (p Point) ScaleBy(factor float64) (sp Point) {
	return Point{p.X * factor, p.Y * factor}
}
