package geometry

func Dot(v0 [3]float64, v1 [3]float64) float64 {
	return v0[0]*v1[0] + v0[1]*v1[1] + v0[2]*v1[2]
}

func Add(v0 [3]float64, v1 [3]float64) [3]float64 {
	return [3]float64{
		v0[0] + v1[0],
		v0[1] + v1[1],
		v0[2] + v1[2],
	}
}

func Subtract(v0 [3]float64, v1 [3]float64) [3]float64 {
	return [3]float64{
		v0[0] - v1[0],
		v0[1] - v1[1],
		v0[2] - v1[2],
	}
}

func At(v [3]float64, t float64) [3]float64 {
	return [3]float64{
		v[0] * t,
		v[1] * t,
		v[2] * t,
	}
}
