package coloured

import (
	"TGPL-exercise-solutions/chapter6/geometry"
	"image/color"
)

type ColouredPoint struct {
	geometry.Point
	Color color.RGBA
}
