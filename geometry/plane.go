package geometry

type Plane struct {
	D float64
	N [3]float64
}

func NewPlane(n [3]float64, p [3]float64) *Plane {
	return &Plane{
		D: Dot(n, p),
		N: n,
	}
}

func (p *Plane) GetIntersectionTriangle(t *Triangle) ([]float64, [][3]float64) {
	tValues := make([]float64, 0)
	intersections := make([][3]float64, 0)

	e0 := Subtract(t.V[1], t.V[0])
	e1 := Subtract(t.V[2], t.V[1])
	e2 := Subtract(t.V[0], t.V[2])

	t0 := (p.D - Dot(p.N, t.V[0])) / Dot(p.N, e0)
	t1 := (p.D - Dot(p.N, t.V[1])) / Dot(p.N, e1)
	t2 := (p.D - Dot(p.N, t.V[2])) / Dot(p.N, e2)

	if t0 >= 0 && t0 <= 1 {
		tValues = append(tValues, t0)
		intersections = append(intersections, Add(t.V[0], At(e0, t0)))
	}

	if t1 >= 0 && t1 <= 1 {
		tValues = append(tValues, t1)
		intersections = append(intersections, Add(t.V[1], At(e1, t1)))
	}

	if t2 >= 0 && t2 <= 1 {
		tValues = append(tValues, t2)
		intersections = append(intersections, Add(t.V[2], At(e2, t2)))
	}

	return tValues, intersections
}
