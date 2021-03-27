package phys

import "github.com/gravitysim/gravity-commons/math"

type BodyLink struct {
	body1 *Body
	body2 *Body

	minDistance float64
	distance    float64

	kappa1    float64
	kappa2    float64
	massRatio float64

	collisionTypeVal collisionType
}

func NewBodyLink(body1, body2 *Body) BodyLink {
	bodyLink := BodyLink{body1: body1, body2: body2}

	bodyLink.minDistance = body1.radius + body2.radius

	tmp := COR / (body1.mass + body2.mass)
	bodyLink.kappa1 = tmp * body2.mass
	bodyLink.kappa2 = tmp * body1.mass
	bodyLink.massRatio = body1.mass / body2.mass

	bodyLink.collisionTypeVal = noCollision

	return bodyLink
}

/**
 * Is bodies in collision now (can lasts multiple steps in sequence)
 */
func (r *BodyLink) IsInCollision() bool {
	r.distance = math.Distance(&r.body1.Position, &r.body2.Position)

	if r.distance < r.minDistance {
		if r.collisionTypeVal == noCollision {
			r.collisionTypeVal = newCollision
		} else {
			r.collisionTypeVal = oldCollision
		}

		return true
	}

	r.collisionTypeVal = noCollision
	return false
}

/**
 * Is collision starts in this step
 */
func (r *BodyLink) IsNewCollision() bool {
	return r.collisionTypeVal == newCollision
}

func (r *BodyLink) GetBody1() *Body {
	return r.body1
}

func (r *BodyLink) GetBody2() *Body {
	return r.body2
}

func (r *BodyLink) GetDistance() float64 {
	return r.distance
}

func (r *BodyLink) GetKappa1() float64 {
	return r.kappa1
}

func (r *BodyLink) GetKappa2() float64 {
	return r.kappa2
}

func (r *BodyLink) GetMassRatio() float64 {
	return r.massRatio
}
