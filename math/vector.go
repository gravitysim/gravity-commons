package math

import "math"

type Vector struct {
	X float64
	Y float64
}

func (r *Vector) Add(x, y float64) *Vector {
	r.X += x
	r.Y += y

	return r
}

func (r *Vector) AddVector(v *Vector) *Vector {
	r.X += v.X
	r.Y += v.Y

	return r
}

func (r *Vector) Sub(x, y float64) *Vector {
	r.X -= x
	r.Y -= y

	return r
}

func (r *Vector) SubVector(v *Vector) *Vector {
	r.X -= v.X
	r.Y -= v.Y

	return r
}

func (r *Vector) Set(x, y float64) *Vector {
	r.X = x
	r.Y = y

	return r
}

func (r *Vector) Clear() *Vector {
	r.Set(0, 0)

	return r
}

func (r *Vector) Length() float64 {
	return math.Sqrt(r.LengthSq())
}

func (r *Vector) LengthSq() float64 {
	return r.X*r.X + r.Y*r.Y
}
