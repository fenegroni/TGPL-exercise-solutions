package geometry

import "math"

type Point struct{ X, Y float64 }

// Distance is a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
