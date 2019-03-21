package main

import (
	"github.com/dkkempto/freed/parser/stl"
	"github.com/dkkempto/freed/slicer"
)

func main() {
	parser := stl.STLParser{}

	mesh := parser.ParseBinary("C:/development/personal/freed/example/dragon.stl")

	s := slicer.NewSlicer(10, 10, 10, 1, 1, 1)
	s.SliceMesh(mesh, slicer.X)
	// fmt.Println(model.String())
}
