package geometry

import "fmt"

const (
	EPSILON = 0.0000001
)

type Triangle struct {
	V [][3]float64
	N [3]float64
}

func (t Triangle) String() string {
	res := fmt.Sprintf("N: (%f, %f, %f)\n", t.N[0], t.N[1], t.N[2])
	res += fmt.Sprintf("V0: (%f, %f, %f)\n", t.V[0][0], t.V[0][1], t.V[0][2])
	res += fmt.Sprintf("V0: (%f, %f, %f)\n", t.V[1][0], t.V[1][1], t.V[1][2])
	res += fmt.Sprintf("V0: (%f, %f, %f)\n", t.V[2][0], t.V[2][1], t.V[2][2])
	return res
}

func (tri *Triangle) GetIntersection(r *Ray) (float64, [3]float64) {
	edge1 := Subtract(tri.V[1], tri.V[0])
	edge2 := Subtract(tri.V[2], tri.V[0])

	h := Cross(r.Dir, edge2)
	a := Dot(edge1, h)

	if a > -EPSILON && a < EPSILON {
		return 0.0, [3]float64{}
	}

	f := 1.0 / a
	s := Subtract(r.Origin, tri.V[0])
	u := f * (Dot(s, h))
	if u < 0.0 || u > 1.0 {
		return 0.0, [3]float64{}
	}
	q := Cross(s, edge1)
	v := f * Dot(r.Dir, q)
	if v < 0.0 || u+v > 1.0 {
		return 0.0, [3]float64{}
	}

	t := f * Dot(edge2, q)
	if t > EPSILON {
		return t, Add(r.Origin, Scale(r.Dir, t))
	}

	return 0.0, [3]float64{}
}

func (t *Triangle) GetBoundingBox() *BoundingBox {
	min := t.V[0]
	max := t.V[0]

	for _, v := range t.V {
		if v[0] < min[0] {
			min[0] = v[0]
		} else if v[0] > max[0] {
			max[0] = v[0]
		}

		if v[1] < min[1] {
			min[1] = v[1]
		} else if v[1] > max[1] {
			max[1] = v[1]
		}

		if v[2] < min[2] {
			min[2] = v[2]
		} else if v[2] > min[2] {
			max[2] = v[2]
		}
	}

	return &BoundingBox{
		Min: min,
		Max: max,
	}
}

func (t *Triangle) GetMidpoint() [3]float64 {
	return [3]float64{
		(t.V[0][0] + t.V[1][0] + t.V[2][0]) / 3,
		(t.V[0][1] + t.V[1][1] + t.V[2][1]) / 3,
		(t.V[0][2] + t.V[1][2] + t.V[2][2]) / 3,
	}
}
