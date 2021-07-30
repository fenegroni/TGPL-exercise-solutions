package main

import (
	"TGPL-exercise-solutions/chapter6/geometry"
	"TGPL-exercise-solutions/chapter6/geometry/coloured"
	"fmt"
	"net/url"
)

func main() {
	p := geometry.Point{X: 0, Y: 0}
	q := geometry.Point{X: 1, Y: 0}
	fmt.Printf("The distance between p %v and q %v is %v\n", p, q, p.Distance(q))
	r := geometry.Point{3, 4}
	fmt.Printf("Point %v scaled by factor %g is %v\n", r, 2.0, r.ScaleBy(2))
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("The perimeter of triangle %v is %v\n", perim, perim.Distance())

	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
	m = nil
	fmt.Println(m["item"])
	// m.Add("item", "3")
	var cp coloured.ColouredPoint
	fmt.Printf("cp %v\n", cp)
}
