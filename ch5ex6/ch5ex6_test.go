package ch5ex6

import (
	"fmt"
	"math"
	"testing"
)

// TestCornerReturnsNaNOrInf reports for what combinations of i and j
// Corner returns an invalid number (NaN or +/-Inf).
func TestCornerReturnsNaNOrInf(t *testing.T) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := Corner(i, j)
			if math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0) {
				t.Logf("Corner(%d, %d) = (%g,%g)", i, j, x, y)
			}
		}
	}
}

// Test_corner is an opaque-box test of Corner vs. OldCorner.
// Given the same input, we expect the same output from both functions.
// If OldCorner returns NaN, we skip to the next iteration so that
// a new Corner implementation can resolve the issue.
func TestCorner(t *testing.T) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := Corner(i, j)
			oldx, oldy := OldCorner(i, j)
			if x != oldx || y != oldy {
				if math.IsNaN(oldx) || math.IsNaN(oldy) {
					continue
				}
				t.Fatalf("Corner(%d, %d) = (%g,%g), want (%g,%g)",
					i, j, x, y, oldx, oldy)
			}
		}
	}
}

func ExampleCorner() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := Corner(i+1, j)
			bx, by := Corner(i, j)
			cx, cy := Corner(i, j+1)
			dx, dy := Corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

// TODO Benchmark Corner vs. OldCorner
