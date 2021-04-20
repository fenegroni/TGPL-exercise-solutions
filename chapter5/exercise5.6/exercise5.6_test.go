package main

import (
	"math"
	"testing"
)

// TestCornerReturnsNaNOrInf reports for what combinations of i and j
// corner returns an invalid number (NaN or +/-Inf).
func TestCornerReturnsNaNOrInf(t *testing.T) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := corner(i, j)
			if math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0) {
				t.Logf("corner(%d, %d) = (%g,%g)", i, j, x, y)
			}
		}
	}
}

// Test_corner is an opaque-box test of corner vs. oldcorner.
// Given the same input, we expect the same output from both functions.
// If oldcorner returns NaN, we skip to the next iteration so that
// a new corner implementation can resolve the issue.
func Test_corner(t *testing.T) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := corner(i, j)
			oldx, oldy := oldcorner(i, j)
			if x != oldx || y != oldy {
				if math.IsNaN(oldx) || math.IsNaN(oldy) {
					continue
				}
				t.Fatalf("corner(%d, %d) = (%g,%g), want (%g,%g)",
					i, j, x, y, oldx, oldy)
			}
		}
	}
}

//TODO Benchmark corner vs. oldcorner
