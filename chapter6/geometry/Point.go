package geometry

import "math"

type Point struct{ X, Y float64 }

// Distance between Points p and q
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy scales the point coordinates by factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
