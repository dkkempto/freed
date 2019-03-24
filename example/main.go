package main

import (
	"fmt"
	"time"

	"github.com/dkkempto/freed/parser/stl"
	"github.com/dkkempto/freed/renderer"
)

func main() {
	parser := stl.STLParser{}

	mesh := parser.ParseBinary("C:/development/personal/freed/example/steps.stl")

	fmt.Println(len(mesh.KDTree.Triangles))
	fmt.Println(len(mesh.KDTree.Left.Triangles))
	fmt.Println(len(mesh.KDTree.Right.Triangles))

	start := time.Now()

	camera := renderer.NewCamera([3]float64{50, 0, 0}, [3]float64{-1, 0, 0}, 100.0, [2]float64{100, 100}, [2]int{200, 200})

	tmp := renderer.Renderer{
		Camera:       camera,
		PathToOutput: "C:/development/personal/freed/res",
	}

	tmp.Render(mesh)
	// s := slicer.NewSlicer(100, 100, 100, 1, 1, 1)
	// s.SliceMesh(mesh, -slicer.X)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed.Seconds())
}
