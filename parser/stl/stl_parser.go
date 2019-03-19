package stl

import (
	"bufio"
	"os"
	"strconv"

	"github.com/dkkempto/freed/geometry"
)

type STLParser struct {
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
				Triangles: make([]geometry.Triangle, 0),
			})
		case "facet":
			models[len(models)-1].Triangles = append(models[len(models)-1].Triangles, geometry.Triangle{
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
		model.BoundingBox = res
	}

	return models
}
