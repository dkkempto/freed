package slicer

import (
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
	var primaryLocation float64
	var primaryIndex int
	var secondaryDimension float64
	var secondaryLocation float64
	var secondaryIndex int
	var stride float64

	switch direction {
	case -X:
		numSlices = s.resX
		dir = [3]float64{-1, 0, 0}
		startingPoint = mesh.BoundingBox.Max

		primaryDimension = -dimensions[1]
		primaryLocation = mesh.BoundingBox.Max[1]
		primaryIndex = 1

		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resX)
	case X:
		numSlices = s.resX
		dir = [3]float64{1, 0, 0}
		startingPoint = mesh.BoundingBox.Min

		primaryDimension = dimensions[1]
		primaryLocation = mesh.BoundingBox.Min[1]
		primaryIndex = 1

		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resX)
	case -Y:
		numSlices = s.resY
		dir = [3]float64{0, -1, 0}
		startingPoint = mesh.BoundingBox.Max

		primaryDimension = dimensions[0]
		primaryLocation = mesh.BoundingBox.Min[0]
		primaryIndex = 0

		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resY)
	case Y:
		numSlices = s.resY
		dir = [3]float64{0, 1, 0}
		startingPoint = mesh.BoundingBox.Min

		primaryDimension = -dimensions[0]
		primaryLocation = mesh.BoundingBox.Max[0]
		primaryIndex = 0

		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resY)

	case -Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, -1}
		startingPoint = mesh.BoundingBox.Max

		primaryDimension = -dimensions[0]
		primaryLocation = mesh.BoundingBox.Max[0]
		primaryIndex = 0

		secondaryDimension = -dimensions[1]
		secondaryLocation = mesh.BoundingBox.Max[1]
		secondaryIndex = 1

		stride = primaryDimension / float64(s.resZ)
	case Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, 1}
		startingPoint = mesh.BoundingBox.Min

		primaryDimension = dimensions[0]
		primaryLocation = mesh.BoundingBox.Min[0]
		primaryIndex = 0

		secondaryDimension = dimensions[1]
		secondaryLocation = -mesh.BoundingBox.Max[1]
		secondaryIndex = 1

		stride = primaryDimension / float64(s.resZ)
	}

	//We want to step through each plane and get the intersection with the meshes
	slices := make([]*Slice, numSlices)

	currentPoint = startingPoint

	// stride :=
	dDir := geometry.At(dir, stride)
	for i := 0; i < numSlices; i++ {
		slice := &Slice{
			Values: make([][]uint8, s.resX),
		}

		for i := range slice.Values {
			slice.Values[i] = make([]uint8, s.resY)
		}
		plane := geometry.NewPlane(dir, currentPoint)

		for _, tri := range mesh.Triangles {
			_, intersections := plane.GetIntersectionTriangle(tri)
			// if len(intersections) > 0 {
			// 	fmt.Println("INTERSECT")
			// }
			for _, intersection := range intersections {
				//0 -> ResX gets mapped to MinBoundingBox -> MaxBoundingBox
				// subtract point from min/divide by dimensions/multiply by resolution :)
				// Screen-coords (x, y) from intersection
				x := int(((intersection[primaryIndex] - primaryLocation) / primaryDimension) * float64(s.resX))
				y := int(((intersection[secondaryIndex] - secondaryLocation) / secondaryDimension) * float64(s.resY))

				slice.Values[x][y] = uint8(1)
				// fmt.Printf("Intersect at: (%v, %v)\n", x, y)
			}
			// if len(intersections) > 0 {
			// 	fmt.Println("ENDINTERSECT")
			// }
		}

		slices = append(slices, slice)
		//Increment current point in dir based on size of bounding box and resolution
		currentPoint = geometry.Add(currentPoint, dDir)
	}

	return slices
}
