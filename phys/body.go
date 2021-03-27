package phys

import "github.com/gravitysim/gravity-commons/math"

type Body struct {
	Position    math.Point
	Velocity    math.Vector
	GravForce   math.Vector
	NormalForce math.Vector

	name        string
	mass        float64
	mu          float64
	radius      float64
	dtMassRatio float64
}

func NewBody(name string, mass, radius float64, position math.Point, velocity math.Vector) Body {
	newBody := Body{name: name, mass: mass, radius: radius, Position: position, Velocity: velocity}
	newBody.GravForce = math.Vector{}
	newBody.NormalForce = math.Vector{}
	newBody.mu = G * mass

	return newBody
}

func (r *Body) GetName() string {
	return r.name
}

func (r *Body) GetMass() float64 {
	return r.mass
}

func (r *Body) GetMu() float64 {
	return r.mu
}

func (r *Body) GetRadius() float64 {
	return r.radius
}

func (r *Body) GetDtMassRatio() float64 {
	return r.dtMassRatio
}

func (r *Body) SetDt(dt float64) {
	r.dtMassRatio = dt / r.mass
}
