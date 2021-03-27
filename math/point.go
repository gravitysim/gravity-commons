package math

import "math"

type Point struct {
	X float64
	Y float64
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(DistanceSq(p1, p2))
}

func DistanceSq(p1, p2 *Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y

	return dx*dx + dy*dy
}
