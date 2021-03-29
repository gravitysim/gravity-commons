package scene_encoding

import (
	"bufio"
	"encoding/binary"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
)

func Encode(scene *scene.GravityScene, writer *bufio.Writer) error {
	return doEncode(createHeader(), scene, writer)
}

func doEncode(header *header, scene *scene.GravityScene, writer *bufio.Writer) error {
	err := encodeHeader(writer, header)
	if err != nil {
		return err
	}

	err = encodeBodies(writer, scene.GetBodies())
	if err != nil {
		return err
	}

	err = encodeLinks(writer, scene)
	if err != nil {
		return err
	}

	err = binary.Write(writer, endian, scene.Step)

	return err
}

func encodeHeader(writer *bufio.Writer, header *header) error {
	err := encodeString(writer, header.name)
	if err != nil {
		return err
	}

	return encodeString(writer, header.version)
}

func encodeBodies(writer *bufio.Writer, bodies []*phys.Body) error {
	err := binary.Write(writer, endian, uint64(len(bodies)))
	if err != nil {
		return err
	}

	for _, body := range bodies {
		err = encodeBody(writer, body)
		if err != nil {
			return err
		}
	}

	return nil
}

func encodeLinks(writer *bufio.Writer, scene *scene.GravityScene) error {
	links := scene.GetLinks()
	err := binary.Write(writer, endian, uint64(len(links)))
	if err != nil {
		return err
	}

	for _, link := range links {
		err = encodeLink(writer, scene, &link)
		if err != nil {
			return err
		}
	}

	return nil
}

func encodeBody(writer *bufio.Writer, body *phys.Body) error {
	err := encodeString(writer, body.GetName())
	if err != nil {
		return err
	}

	err = binary.Write(writer, endian, body.GetMass())
	if err != nil {
		return err
	}

	err = binary.Write(writer, endian, body.GetRadius())
	if err != nil {
		return err
	}

	err = binary.Write(writer, endian, body.Position)
	if err != nil {
		return err
	}

	return binary.Write(writer, endian, body.Velocity)
}

func encodeLink(writer *bufio.Writer, scene *scene.GravityScene, link *phys.BodyLink) error {
	body1Num := uint64(scene.GetBodyNum(link.GetBody1()))
	body2Num := uint64(scene.GetBodyNum(link.GetBody2()))

	err := binary.Write(writer, endian, body1Num)
	if err != nil {
		return err
	}

	return binary.Write(writer, endian, body2Num)
}

func encodeString(writer *bufio.Writer, value string) error {
	byteValue := []byte(value)

	err := binary.Write(writer, endian, uint8(len(byteValue)))
	if err != nil {
		return err
	}

	_, err = writer.Write(byteValue)

	return err
}
