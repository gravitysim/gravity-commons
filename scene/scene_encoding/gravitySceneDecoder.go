package scene_encoding

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/gravitysim/gravity-commons/math"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
)

func Decode(reader *bufio.Reader) (*scene.GravityScene, error) {
	readHeader, err := decodeHeader(reader)

	err = checkHeader(readHeader)
	if err != nil {
		return nil, err
	}

	gravityScene := scene.GravityScene{}
	err = decodeBodies(reader, &gravityScene)
	if err != nil {
		return nil, err
	}

	err = decodeLinks(reader, &gravityScene)
	if err != nil {
		return nil, err
	}

	var step uint64
	err = binary.Read(reader, endian, &step)
	gravityScene.Step = step

	return &gravityScene, err
}

func checkHeader(header *header) error {
	supportedHeader := createHeader()

	if supportedHeader.name != header.name {
		return errors.New("invalid header: " + header.name)
	}
	if supportedHeader.version != header.version {
		return errors.New("invalid version: " + header.version)
	}

	return nil
}

func decodeHeader(reader *bufio.Reader) (*header, error) {
	headerName, err := decodeString(reader)
	if err != nil {
		return nil, err
	}

	version, err := decodeString(reader)

	return &header{name: headerName, version: version}, nil
}

func decodeBodies(reader *bufio.Reader, scene *scene.GravityScene) error {
	var bodiesCount uint64
	err := binary.Read(reader, endian, &bodiesCount)
	if err != nil {
		return err
	}

	for i := uint64(0); i < bodiesCount; i++ {
		err = decodeBody(reader, scene)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeLinks(reader *bufio.Reader, scene *scene.GravityScene) error {
	var linksCount uint64
	err := binary.Read(reader, endian, &linksCount)
	if err != nil {
		return err
	}

	for i := uint64(0); i < linksCount; i++ {
		err = decodeLink(reader, scene)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeBody(reader *bufio.Reader, scene *scene.GravityScene) error {
	name, err := decodeString(reader)
	if err != nil {
		return err
	}

	var mass float64
	err = binary.Read(reader, endian, &mass)
	if err != nil {
		return err
	}

	var radius float64
	err = binary.Read(reader, endian, &radius)
	if err != nil {
		return err
	}

	var position math.Point
	err = binary.Read(reader, endian, &position)
	if err != nil {
		return err
	}

	var velocity math.Vector
	err = binary.Read(reader, endian, &velocity)
	if err != nil {
		return err
	}

	body := phys.NewBody(name, mass, radius, position, velocity)
	scene.AddBody(&body)

	return nil
}

func decodeLink(reader *bufio.Reader, scene *scene.GravityScene) error {
	var body1Num uint64
	err := binary.Read(reader, endian, &body1Num)
	if err != nil {
		return err
	}

	var body2Num uint64
	err = binary.Read(reader, endian, &body2Num)
	if err != nil {
		return err
	}

	scene.LinkBodies(body1Num, body2Num)
	return nil
}

func decodeString(reader *bufio.Reader) (string, error) {
	var length uint8
	err := binary.Read(reader, endian, &length)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	_, err = reader.Read(strBytes)
	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}
