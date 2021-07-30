package main

import (
	"TGPL-exercise-solutions/chapter6/geometry"
	"fmt"
)

func main() {
	var p, q geometry.Point
	fmt.Printf("The distance between %v and %v is %v", p, q, p.Distance(q))
	fmt.Printf("The distance between %v and %v is %v", p, q, geometry.Distance(p, q))
}
