package geometry

import "fmt"

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
