package geometry

import "fmt"

type Mesh struct {
	Name        string
	BoundingBox *BoundingBox
	KDTree      *KDNode
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

func (m *Mesh) GetIntersections(r *Ray) ([]float64, [][3]float64) {
	//TODO: Implement this
	/**
	This method will take in a ray, and return the intersections with the mesh
	*/
	tValues := make([]float64, 0)
	intersections := make([][3]float64, 0)

	for _, triangle := range m.Triangles {
		t, intersection := triangle.GetIntersection(r)
		if t != 0.0 {
			tValues = append(tValues, t)
			intersections = append(intersections, intersection)
		}
	}

	return tValues, intersections

	// return m.KDTree.GetIntersections(r)
}

func (m *Mesh) String() string {
	res := fmt.Sprintf("Name: %v\n", m.Name)
	res += fmt.Sprintf("%v\n", m.BoundingBox.String())
	return res
}
