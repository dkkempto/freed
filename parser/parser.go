package parser

type Parser interface {
	parse(path string) []Model
}

type Triangle struct {
	V0 [3]float64
	V1 [3]float64
	V2 [3]float64
	N0 [3]float64
	N1 [3]float64
	N2 [3]float64
}

type Model struct {
	Name      string
	Triangles []Triangle
}
