package math

import "testing"

func TestVectorAdd(t *testing.T) {
	v1 := Vector{X: 10, Y: 20}
	v1.Add(5, 6)

	if !(v1.X == 15 && v1.Y == 26) {
		t.Errorf("Invalid add operation\n")
		return
	}

	v2 := Vector{X: 100, Y: 200}
	v1.AddVector(&v2)

	if !(v1.X == 115 && v1.Y == 226 && v2.X == 100 && v2.Y == 200) {
		t.Errorf("Invalid add vector operation\n")
	}
}

func TestVectorSub(t *testing.T) {
	v1 := Vector{X: 10, Y: 20}
	v1.Sub(4, 6)

	if !(v1.X == 6 && v1.Y == 14) {
		t.Errorf("Invalid sub operation\n")
		return
	}

	v2 := Vector{X: 100, Y: 200}
	v1.SubVector(&v2)

	if !(v1.X == -94 && v1.Y == -186 && v2.X == 100 && v2.Y == 200) {
		t.Errorf("Invalid sub vector operation\n")
	}
}

func TestVectorSet(t *testing.T) {
	v1 := Vector{X: 10, Y: 20}
	v1.Set(4, 6)

	if !(v1.X == 4 && v1.Y == 6) {
		t.Errorf("Invalid set operation\n")
	}
}

func TestVectorClear(t *testing.T) {
	v1 := Vector{X: 10, Y: 20}
	v1.Clear()

	if !(v1.X == 0 && v1.Y == 0) {
		t.Errorf("Invalid clear operation\n")
	}
}
