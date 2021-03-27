package math

type Vector struct {
	X float64
	Y float64
}

func (r *Vector) Add(x, y float64) {
	r.X += x
	r.Y += y
}

func (r *Vector) AddVector(v *Vector) {
	r.X += v.X
	r.Y += v.Y
}

func (r *Vector) Sub(x, y float64) {
	r.X -= x
	r.Y -= y
}

func (r *Vector) SubVector(v *Vector) {
	r.X -= v.X
	r.Y -= v.Y
}

func (r *Vector) Set(x, y float64) {
	r.X = x
	r.Y = y
}

func (r *Vector) Clear() {
	r.Set(0, 0)
}
