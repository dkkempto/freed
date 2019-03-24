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

func (bb *BoundingBox) Expand(t *Triangle) {
	for _, v := range t.V {
		if v[0] < bb.Min[0] {
			bb.Min[0] = v[0]
		} else if v[0] > bb.Max[0] {
			bb.Max[0] = v[0]
		}

		if v[1] < bb.Min[1] {
			bb.Min[1] = v[1]
		} else if v[1] > bb.Max[1] {
			bb.Max[1] = v[1]
		}

		if v[2] < bb.Min[2] {
			bb.Min[2] = v[2]
		} else if v[2] > bb.Max[2] {
			bb.Max[2] = v[2]
		}
	}
}

func (bb *BoundingBox) GetLongestAxis() int {
	res := 0
	maxLength := 0.0

	for i := 0; i < 3; i++ {
		l := bb.Max[i] - bb.Min[i]
		if l > maxLength {
			res = i
		}
	}

	return res
}

func (bb *BoundingBox) Intersects(r *Ray) bool {
	//TODO: Optimize this :) Can add inverse calc to ray object, so we aren't doing redundantly here
	invdir := [3]float64{
		1 / r.Dir[0],
		1 / r.Dir[1],
		1 / r.Dir[2],
	}

	tmin := (bb.Min[0] - r.Origin[0]) * invdir[0]
	tmax := (bb.Max[0] - r.Origin[0]) * invdir[0]

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	tymin := (bb.Min[1] - r.Origin[1]) * invdir[1]
	tymax := (bb.Max[1] - r.Origin[1]) * invdir[1]

	if tymin > tymax {
		tymin, tymax = tymax, tymin
	}

	if tmin > tymax || tymin > tmax {
		return false
	}

	if tymax < tmax {
		tmax = tymax
	}

	tzmin := (bb.Min[2] - r.Origin[2]) * invdir[2]
	tzmax := (bb.Max[2] - r.Origin[2]) * invdir[2]

	if tzmin > tzmax {
		tzmin, tzmax = tzmax, tzmin
	}

	if tmin > tzmax || tzmin > tmax {
		return false
	}

	if tzmin > tmin {
		tmin = tzmin
	}

	if tzmax < tmax {
		tmax = tzmax
	}

	return true
}

func (bb *BoundingBox) String() string {
	res := fmt.Sprintf("Min: (%f, %f, %f)\n", bb.Min[0], bb.Min[1], bb.Min[2])
	res += fmt.Sprintf("Max: (%f, %f, %f)\n", bb.Max[0], bb.Max[1], bb.Max[2])
	return res
}
