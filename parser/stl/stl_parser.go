package stl

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
	"strconv"

	"github.com/dkkempto/freed/geometry"
)

type STLParser struct {
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (p STLParser) Parse(path string) []*geometry.Mesh {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	models := make([]*geometry.Mesh, 0)

	for scanner.Scan() {
		token := scanner.Text()

		switch token {
		case "solid":
			scanner.Scan()
			models = append(models, &geometry.Mesh{
				Name:      scanner.Text(),
				Triangles: make([]*geometry.Triangle, 0),
			})
		case "facet":
			models[len(models)-1].Triangles = append(models[len(models)-1].Triangles, &geometry.Triangle{
				V: make([][3]float64, 0),
			})
		case "normal":
			normal := [3]float64{}
			currModel := len(models) - 1
			currTri := len(models[currModel].Triangles) - 1

			for i := 0; i < 3; i++ {
				scanner.Scan()
				normal[i], err = strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					panic(err)
				}
			}

			models[currModel].Triangles[currTri].N = normal
		case "outer":
			continue
		case "loop":
			continue
		case "vertex":
			vertex := [3]float64{}
			currModel := len(models) - 1
			currTri := len(models[currModel].Triangles) - 1

			for i := 0; i < 3; i++ {
				scanner.Scan()
				vertex[i], err = strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					panic(err)
				}
			}

			models[currModel].Triangles[currTri].V = append(models[currModel].Triangles[currTri].V, vertex)
		case "endloop":
		case "endfacet":
		case "endsolid":
		}
	}

	for _, model := range models {
		res := geometry.NewBoundingBox(model)
		model.BoundingBox = &res
	}

	return models
}

func (p STLParser) ParseBinary(path string) *geometry.Mesh {
	f, err := os.Open(path)
	check(err)

	//Skip the header for now. Not sure what is even in it :)
	f.Seek(80, 0)
	bNumTriangles := make([]byte, 4)
	_, err = f.Read(bNumTriangles)
	check(err)

	numTriangles := binary.LittleEndian.Uint32(bNumTriangles)

	model := &geometry.Mesh{
		Triangles: make([]*geometry.Triangle, numTriangles),
	}

	var i uint32

	for i = 0; i < numTriangles; i++ {
		bTriangle := make([]byte, 50)
		_, err = f.Read(bTriangle)
		check(err)
		var n0, n1, n2,
			v0f0, v0f1, v0f2,
			v1f0, v1f1, v1f2,
			v2f0, v2f1, v2f2 float32

		buf := bytes.NewReader(bTriangle[0:48])

		err = binary.Read(buf, binary.LittleEndian, &n0)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &n1)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &n2)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v0f0)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v0f1)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v0f2)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v1f0)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v1f1)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v1f2)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v2f0)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v2f1)
		check(err)
		err = binary.Read(buf, binary.LittleEndian, &v2f2)
		check(err)

		model.AddTriangle(
			&geometry.Triangle{
				N: [3]float64{float64(n0), float64(n1), float64(n2)},
				V: [][3]float64{
					{float64(v0f0), float64(v0f1), float64(v0f2)},
					{float64(v1f0), float64(v1f1), float64(v1f2)},
					{float64(v2f0), float64(v2f1), float64(v2f2)},
				},
			},
		)
	}

	return model

}
