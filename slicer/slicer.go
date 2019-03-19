package slicer

import (
	"github.com/dkkempto/freed/parser"
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

func SliceModels(s *Slicer, models []parser.Model, direction int) []Slice {
	var numSlices int

	switch direction {
	case -X:
		numSlices = s.resX
	case X:
		numSlices = s.resX
	case -Y:
		numSlices = s.resY
	case Y:
		numSlices = s.resY
	case -Z:
		numSlices = s.resZ
	case Z:
		numSlices = s.resZ
	}

	res := make([]Slice, numSlices)

	//We want to step through each plane and get the intersection with the meshes

	for i := 0; i < numSlices; i++ {

	}

	return res
}
