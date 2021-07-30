package geometry

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

// PathDistance returns the distance travelled along the path.
func PathDistance(p Path) float64 {
	return p.Distance()
}
