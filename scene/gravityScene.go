package scene

import (
	"github.com/gravitysim/gravity-commons/phys"
)

type GravityScene struct {
	Step uint64

	bodies []phys.Body
	links  []phys.BodyLink

	bodiesNum map[*phys.Body]int
}

func (r *GravityScene) GetBodies() []phys.Body {
	return r.bodies
}

func (r *GravityScene) GetLinks() []phys.BodyLink {
	return r.links
}

func (r *GravityScene) AddBody(body *phys.Body) {
	if r.bodiesNum == nil {
		r.bodiesNum = make(map[*phys.Body]int)
	}
	r.bodiesNum[body] = len(r.bodies)
	r.bodies = append(r.bodies, *body)
}

func (r *GravityScene) LinkBodies(body1Num, body2Num uint64) {
	r.links = append(r.links, phys.NewBodyLink(&r.bodies[body1Num], &r.bodies[body2Num]))
}

func (r *GravityScene) GetBodyNum(body *phys.Body) int {
	return r.bodiesNum[body]
}
