package phys

type collisionType int

const (
	noCollision collisionType = iota
	newCollision
	oldCollision
)
