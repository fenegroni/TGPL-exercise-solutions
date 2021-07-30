package main

import (
	"TGPL-exercise-solutions/chapter6/geometry"
	"fmt"
)

func main() {
	p := geometry.Point{X: 0, Y: 0}
	q := geometry.Point{X: 1, Y: 0}
	fmt.Printf("The distance between %v and %v is %v\n", p, q, p.Distance(q))
	var path geometry.Path
	r := geometry.Point{X: 0, Y: 1}
	path = append(path, p, q, r)
	fmt.Printf("The distance on path %v is %v\n", path, path.Distance())
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("The perimeter of triangle %v is %v\n", perim, perim.Distance())
}
