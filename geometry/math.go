package geometry

import (
	"math"
)

var (
	GLOBAL_UP = [3]float64{0, 0, 1}
)

func Dot(v0 [3]float64, v1 [3]float64) float64 {
	return v0[0]*v1[0] + v0[1]*v1[1] + v0[2]*v1[2]
}

func Cross(u [3]float64, v [3]float64) [3]float64 {
	return [3]float64{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
}

func Add(vectors ...[3]float64) [3]float64 {
	res := [3]float64{0, 0, 0}
	for _, v := range vectors {
		res[0] += v[0]
		res[1] += v[1]
		res[2] += v[2]
	}
	return res
}

func Subtract(v0 [3]float64, v1 [3]float64) [3]float64 {
	return [3]float64{
		v0[0] - v1[0],
		v0[1] - v1[1],
		v0[2] - v1[2],
	}
}

func Scale(v [3]float64, t float64) [3]float64 {
	return [3]float64{
		v[0] * t,
		v[1] * t,
		v[2] * t,
	}
}

func Normalize(v [3]float64) [3]float64 {
	length := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	res := v
	if length > 0 {
		res = [3]float64{
			v[0] / length,
			v[1] / length,
			v[2] / length,
		}
	}
	return res
}
