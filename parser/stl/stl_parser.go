package stl

import (
	"bufio"
	"os"
	"strconv"

	"github.com/dkkempto/freed/parser"
)

type STLParser struct {
}

func (p STLParser) Parse(path string) []parser.Model {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	models := make([]parser.Model, 0)

	for scanner.Scan() {
		token := scanner.Text()

		switch token {
		case "solid":
			scanner.Scan()
			models = append(models, parser.Model{
				Name:      scanner.Text(),
				Triangles: make([]parser.Triangle, 0),
			})
		case "facet":
			models[len(models)-1].Triangles = append(models[len(models)-1].Triangles, parser.Triangle{
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

	return models
}
