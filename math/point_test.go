package math

import "testing"

func TestPointAdd(t *testing.T) {
	p := Point{X: 10, Y: 20}
	p.Add(5, 6)

	if !(p.X == 15 && p.Y == 26) {
		t.Errorf("Invalid add operation\n")
	}
}

func TestPointSub(t *testing.T) {
	p := Point{X: 10, Y: 20}
	p.Sub(4, 6)

	if !(p.X == 6 && p.Y == 14) {
		t.Errorf("Invalid sub operation\n")
	}
}

func TestPointSet(t *testing.T) {
	p := Point{X: 10, Y: 20}
	p.Set(4, 6)

	if !(p.X == 4 && p.Y == 6) {
		t.Errorf("Invalid set operation\n")
	}
}

func TestPointClear(t *testing.T) {
	p := Point{X: 10, Y: 20}
	p.Clear()

	if !(p.X == 0 && p.Y == 0) {
		t.Errorf("Invalid clear operation\n")
	}
}
