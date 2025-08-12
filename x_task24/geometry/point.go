package geometry

import "math"

type Point struct {
	x, y float64
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(
		math.Pow((other.x-p.x), 2) + math.Pow((other.y-p.y), 2),
	)
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}
