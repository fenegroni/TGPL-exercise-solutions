package geometry

import "math"

type Point struct{ X, Y float64 }

// Distance is a traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance is a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// The Distance travelled along the path.
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
