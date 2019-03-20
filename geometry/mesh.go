package geometry

import "fmt"

type Mesh struct {
	Name        string
	BoundingBox *BoundingBox
	Triangles   []*Triangle
}

func (m *Mesh) AddTriangle(t *Triangle) {
	if m.BoundingBox == nil {
		m.BoundingBox = &BoundingBox{
			Min: t.V[0],
			Max: t.V[0],
		}
	}

	for _, v := range t.V {
		for i := 0; i < 3; i++ {
			if v[i] < m.BoundingBox.Min[i] {
				m.BoundingBox.Min[i] = v[i]
			}

			if v[i] > m.BoundingBox.Max[i] {
				m.BoundingBox.Max[i] = v[i]
			}
		}
	}

	m.Triangles = append(m.Triangles, t)
}

func (m *Mesh) String() string {
	res := fmt.Sprintf("Name: %v\n", m.Name)
	// for _, t := range m.Triangles {
	// 	res += fmt.Sprintf("%v\n", t.String())
	// }
	res += fmt.Sprintf("%v\n", m.BoundingBox.String())
	return res
}
