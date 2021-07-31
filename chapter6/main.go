package main

import (
	"TGPL-exercise-solutions/chapter6/geometry"
	"TGPL-exercise-solutions/chapter6/geometry/coloured"
	"fmt"
	"image/color"
	"net/url"
	"sync"
)

type AType struct {
	A int
}

type BType struct {
	B int
}

func (at *AType) Increment() {
	at.A++
}

func (bt *BType) Increment() {
	bt.B++
}

type ABType struct {
	AType
	BType
}

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	at := ABType{AType{1}, BType{2}}
	at.AType.Increment()

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
	cp := coloured.ColouredPoint{Point: &geometry.Point{X: 1, Y: 2}, RGBA: &color.RGBA{A: 1, B: 2, G: 3, R: 4}}
	fmt.Printf("cp %v\n", cp)
	x := cp.ScaleBy(3.1)
	fmt.Printf("x: %v\n", x)

	mars := geometry.Point{1, 2}
	pluto := geometry.Point{2, 3}
	distanceFromMars := mars.Distance
	fmt.Println(distanceFromMars(pluto))
}
