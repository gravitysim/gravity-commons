package scene_encoding

import (
	"bufio"
	"bytes"
	"github.com/gravitysim/gravity-commons/math"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
	"reflect"
	"testing"
)

func TestEncoding(t *testing.T) {
	gravityScene := createGravityScene()

	buffer := new(bytes.Buffer)
	writer := bufio.NewWriter(buffer)
	err := Encode(gravityScene, writer)
	if err != nil {
		t.Errorf("Error while encode scene: %s\n", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		t.Errorf("Error in flushing scene: %s\n", err)
		return
	}

	var decodedScene *scene.GravityScene

	buffer = bytes.NewBuffer(buffer.Bytes())
	reader := bufio.NewReader(buffer)
	decodedScene, err = Decode(reader)
	if err != nil {
		t.Errorf("Error while decode scene: %s\n", err)
		return
	}

	assertEquals(t, gravityScene, decodedScene)
}

func createGravityScene() *scene.GravityScene {
	gravityScene := scene.GravityScene{Step: 100}
	b1 := phys.NewBody("b1", 1, 2, math.Point{X: 3, Y: 4}, math.Vector{X: 5, Y: 6})
	b2 := phys.NewBody("body2", 7, 8, math.Point{X: 9, Y: 10}, math.Vector{X: 11, Y: 12})
	b3 := phys.NewBody("longNameBody3", 13, 14, math.Point{X: 15, Y: 16}, math.Vector{X: 17, Y: 18})
	gravityScene.AddBody(&b1)
	gravityScene.AddBody(&b2)
	gravityScene.AddBody(&b3)

	gravityScene.LinkBodies(0, 1)
	gravityScene.LinkBodies(0, 2)

	return &gravityScene
}

func assertEquals(t *testing.T, expected *scene.GravityScene, actual *scene.GravityScene) {
	res := reflect.DeepEqual(expected.GetBodies(), actual.GetBodies())
	res = res && reflect.DeepEqual(expected.GetLinks(), actual.GetLinks())
	res = res && reflect.DeepEqual(expected.Step, actual.Step)

	if !res {
		t.Errorf("Could not restore original state!")
	}
}
