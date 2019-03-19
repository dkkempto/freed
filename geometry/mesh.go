package geometry

import "fmt"

type Mesh struct {
	Name        string
	BoundingBox BoundingBox
	Triangles   []Triangle
}

func (m Mesh) String() string {
	res := fmt.Sprintf("Name: %v\n", m.Name)
	for _, t := range m.Triangles {
		res += fmt.Sprintf("%v\n", t.String())
	}
	res += fmt.Sprintf("%v\n", m.BoundingBox.String())
	return res
}
