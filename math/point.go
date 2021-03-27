package math

import "math"

type Point struct {
	X float64
	Y float64
}

func (r *Point) Add(x, y float64) {
	r.X += x
	r.Y += y
}

func (r *Point) Sub(x, y float64) {
	r.X -= x
	r.Y -= y
}

func (r *Point) Set(x, y float64) {
	r.X = x
	r.Y = y
}

func (r *Point) Clear() {
	r.Set(0, 0)
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(DistanceSq(p1, p2))
}

func DistanceSq(p1, p2 *Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y

	return dx*dx + dy*dy
}
