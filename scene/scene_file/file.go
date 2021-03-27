package scene_file

import (
	"bufio"
	"github.com/gravitysim/gravity-commons/scene"
	"github.com/gravitysim/gravity-commons/scene/scene_encoding"
	"os"
)

func LoadScene(fileName string) (*scene.GravityScene, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	return scene_encoding.Decode(reader)
}

func SaveScene(scene *scene.GravityScene, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	err = scene_encoding.Encode(scene, writer)
	if err != nil {
		return err
	}

	return writer.Flush()
}
