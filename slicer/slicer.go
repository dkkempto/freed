package slicer

import (
	"fmt"

	"github.com/dkkempto/freed/geometry"
)

const (
	_ = iota
	X = iota
	Y = iota
	Z = iota
)

type Slice struct {
	Values [][]uint8
}

type Plane struct {
}

type Slicer struct {
	resX, resY, resZ int
	dx, dy, dz       float64
}

func NewSlicer(resX int, resY int, resZ int, dx float64, dy float64, dz float64) Slicer {
	return Slicer{
		resX: resX,
		resY: resY,
		resZ: resZ,
		dx:   dx,
		dy:   dy,
		dz:   dz,
	}
}

func (s *Slicer) SliceMesh(mesh *geometry.Mesh, direction int) []*Slice {
	var numSlices int
	var dir [3]float64
	var startingPoint [3]float64
	var currentPoint [3]float64
	dimensions := geometry.Subtract(mesh.BoundingBox.Max, mesh.BoundingBox.Min)
	var primaryDimension float64
	var stride float64

	switch direction {
	case -X:
		numSlices = s.resX
		dir = [3]float64{-1, 0, 0}
		startingPoint = mesh.BoundingBox.Min
		primaryDimension = dimensions[0]
		stride = primaryDimension / float64(s.resX)
	case X:
		numSlices = s.resX
		dir = [3]float64{1, 0, 0}
		startingPoint = mesh.BoundingBox.Max
		primaryDimension = dimensions[0]
		stride = primaryDimension / float64(s.resX)
	case -Y:
		numSlices = s.resY
		dir = [3]float64{0, -1, 0}
		startingPoint = mesh.BoundingBox.Min
		primaryDimension = dimensions[1]
		stride = primaryDimension / float64(s.resY)
	case Y:
		numSlices = s.resY
		dir = [3]float64{0, 1, 0}
		startingPoint = mesh.BoundingBox.Max
		primaryDimension = dimensions[1]
		stride = primaryDimension / float64(s.resY)

	case -Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, -1}
		startingPoint = mesh.BoundingBox.Min
		primaryDimension = dimensions[2]
		stride = primaryDimension / float64(s.resZ)
	case Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, 1}
		startingPoint = mesh.BoundingBox.Max
		primaryDimension = dimensions[2]
		stride = primaryDimension / float64(s.resZ)
	}

	//We want to step through each plane and get the intersection with the meshes
	slices := make([]*Slice, numSlices)

	currentPoint = startingPoint

	// stride :=
	dDir := geometry.At(dir, stride)
	for i := 0; i < numSlices; i++ {
		slice := Slice{}
		plane := geometry.NewPlane(dir, currentPoint)

		for _, tri := range mesh.Triangles {
			_, intersections := plane.GetIntersectionTriangle(tri)
			for _, intersection := range intersections {
				//0 -> ResX gets mapped to MinBoundingBox -> MaxBoundingBox
				// screenX :=
				fmt.Println(intersection)
			}
		}

		slices = append(slices, &slice)
		//Increment current point in dir based on size of bounding box and resolution
		currentPoint = geometry.Add(currentPoint, dDir)
	}

	return slices
}
