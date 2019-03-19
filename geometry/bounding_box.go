package geometry

import "fmt"

type BoundingBox struct {
	Min [3]float64
	Max [3]float64
}

func NewBoundingBox(mesh *Mesh) BoundingBox {
	min := [3]float64{0, 0, 0}
	max := [3]float64{0, 0, 0}

	for tIndex, tri := range mesh.Triangles {
		if tIndex == 0 {
			min = tri.V[0]
			max = tri.V[0]
		}

		for _, v := range tri.V {
			if v[0] < min[0] {
				min[0] = v[0]
			}

			if v[1] < min[1] {
				min[1] = v[1]
			}

			if v[2] < min[2] {
				min[2] = v[2]
			}

			if v[0] > max[0] {
				max[0] = v[0]
			}

			if v[1] > max[1] {
				max[1] = v[1]
			}

			if v[2] > max[2] {
				max[2] = v[2]
			}
		}
	}

	res := BoundingBox{
		Min: min,
		Max: max,
	}

	return res
}

func (bb *BoundingBox) String() string {
	res := fmt.Sprintf("Min: (%f, %f, %f)\n", bb.Min[0], bb.Min[1], bb.Min[2])
	res += fmt.Sprintf("Max: (%f, %f, %f)\n", bb.Max[0], bb.Max[1], bb.Max[2])
	return res
}
