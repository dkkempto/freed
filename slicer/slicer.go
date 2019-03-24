package slicer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

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
	var width, height int

	switch direction {
	case -X:
		numSlices = s.resX
		dir = [3]float64{-1, 0, 0}
		startingPoint = mesh.BoundingBox.Max

		width = s.resY
		primaryDimension = -dimensions[1]
		primaryLocation = mesh.BoundingBox.Max[1]
		primaryIndex = 1

		height = s.resZ
		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resX)
	case X:
		numSlices = s.resX
		dir = [3]float64{1, 0, 0}
		startingPoint = mesh.BoundingBox.Min

		width = s.resY
		primaryDimension = dimensions[1]
		primaryLocation = mesh.BoundingBox.Min[1]
		primaryIndex = 1

		height = s.resZ
		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resX)
	case -Y:
		numSlices = s.resY
		dir = [3]float64{0, -1, 0}
		startingPoint = mesh.BoundingBox.Max

		width = s.resX
		primaryDimension = dimensions[0]
		primaryLocation = mesh.BoundingBox.Min[0]
		primaryIndex = 0

		height = s.resZ
		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resY)
	case Y:
		numSlices = s.resY
		dir = [3]float64{0, 1, 0}
		startingPoint = mesh.BoundingBox.Min

		width = s.resX
		primaryDimension = -dimensions[0]
		primaryLocation = mesh.BoundingBox.Max[0]
		primaryIndex = 0

		height = s.resZ
		secondaryDimension = -dimensions[2]
		secondaryLocation = mesh.BoundingBox.Max[2]
		secondaryIndex = 2

		stride = primaryDimension / float64(s.resY)

	case -Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, -1}
		startingPoint = mesh.BoundingBox.Max

		width = s.resX
		primaryDimension = -dimensions[0]
		primaryLocation = mesh.BoundingBox.Max[0]
		primaryIndex = 0

		height = s.resY
		secondaryDimension = -dimensions[1]
		secondaryLocation = mesh.BoundingBox.Max[1]
		secondaryIndex = 1

		stride = primaryDimension / float64(s.resZ)
	case Z:
		numSlices = s.resZ
		dir = [3]float64{0, 0, 1}
		startingPoint = mesh.BoundingBox.Min

		width = s.resX
		primaryDimension = dimensions[0]
		primaryLocation = mesh.BoundingBox.Min[0]
		primaryIndex = 0

		height = s.resY
		secondaryDimension = dimensions[1]
		secondaryLocation = -mesh.BoundingBox.Max[1]
		secondaryIndex = 1

		stride = primaryDimension / float64(s.resZ)
	}

	//We want to step through each plane and get the intersection with the meshes
	slices := make([]*Slice, numSlices)

	currentPoint = startingPoint

	// stride :=
	dDir := geometry.Scale(dir, stride)
	for i := 0; i < numSlices; i++ {
		slice := &Slice{
			Values: make([][]uint8, width),
		}

		upLeft := image.Point{0, 0}
		lowRight := image.Point{width, height}

		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		cyan := color.RGBA{100, 200, 200, 0xff}

		for j := range slice.Values {
			slice.Values[j] = make([]uint8, height)
		}

		plane := geometry.NewPlane(dir, currentPoint)

		for _, tri := range mesh.Triangles {
			_, intersections := plane.GetIntersectionTriangle(tri)
			for _, intersection := range intersections {
				x := int(((intersection[primaryIndex] - primaryLocation) / primaryDimension) * float64(width))
				y := int(((intersection[secondaryIndex] - secondaryLocation) / secondaryDimension) * float64(height))

				img.SetRGBA(x, y, cyan)
				slice.Values[x][y] = uint8(1)
			}
		}

		fileName := fmt.Sprintf("C:/development/personal/freed/res/slice_%08d.png", i)
		f, _ := os.Create(fileName)
		png.Encode(f, img)
		f.Close()

		// slices = append(slices, slice)
		//Increment current point in dir based on size of bounding box and resolution
		currentPoint = geometry.Add(currentPoint, dDir)
	}

	return slices
}
