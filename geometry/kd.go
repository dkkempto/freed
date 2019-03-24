package geometry

type KDNode struct {
	BoundingBox *BoundingBox
	Left        *KDNode
	Right       *KDNode
	Triangles   []*Triangle
}

func BuildKDNode(triangles []*Triangle, depth int) *KDNode {
	node := &KDNode{}
	node.Triangles = triangles
	node.Left = nil
	node.Right = nil
	node.BoundingBox = &BoundingBox{}

	if len(triangles) == 0 {
		return node
	}

	if len(triangles) == 1 {
		node.BoundingBox = triangles[0].GetBoundingBox()
		node.Left = &KDNode{}
		node.Right = &KDNode{}
		node.Left.Triangles = make([]*Triangle, 0)
		node.Right.Triangles = make([]*Triangle, 0)
		return node
	}

	node.BoundingBox = triangles[0].GetBoundingBox()

	for _, triangle := range triangles[1:] {
		node.BoundingBox.Expand(triangle)
	}

	midpt := [3]float64{0.0, 0.0, 0.0}

	inv_num_triangles := 1.0 / float64(len(triangles))

	for _, triangle := range triangles {
		midpt = Add(midpt, Scale(triangle.GetMidpoint(), inv_num_triangles))
	}

	left_tris := make([]*Triangle, 0)
	right_tris := make([]*Triangle, 0)
	axis := node.BoundingBox.GetLongestAxis()
	for _, triangle := range triangles {
		if midpt[axis] >= triangle.GetMidpoint()[axis] {
			left_tris = append(left_tris, triangle)
		} else {
			right_tris = append(right_tris, triangle)
		}
	}

	if len(left_tris) == 0 && len(right_tris) > 0 {
		left_tris = right_tris
	}

	if len(right_tris) == 0 && len(left_tris) > 0 {
		right_tris = left_tris
	}

	matches := 0
	for _, left_tri := range left_tris {
		for _, right_tri := range right_tris {
			if left_tri == right_tri {
				matches++
			}
		}
	}

	if float64(matches)/float64(len(left_tris)) < 0.5 && float64(matches)/float64(len(right_tris)) < 0.5 {
		node.Left = BuildKDNode(left_tris, depth+1)
		node.Right = BuildKDNode(right_tris, depth+1)
	} else {
		node.Left = &KDNode{}
		node.Right = &KDNode{}
		node.Left.Triangles = make([]*Triangle, 0)
		node.Right.Triangles = make([]*Triangle, 0)
	}

	return node
}

func (node *KDNode) GetIntersections(r *Ray) ([]float64, [][3]float64) {
	if node.BoundingBox.Intersects(r) {
		// var normal [3]float64
		// hit_tri := false
		// var hit_pt, local_hit_pt [3]float64

		if len(node.Left.Triangles) > 0 || len(node.Right.Triangles) > 0 {
			tValuesLeft, intersectionsLeft := node.Left.GetIntersections(r)
			tValuesRight, intersectionsRight := node.Right.GetIntersections(r)
			return append(tValuesLeft, tValuesRight...), append(intersectionsLeft, intersectionsRight...)
		} else {
			tValues := make([]float64, 0)
			intersections := make([][3]float64, 0)
			for _, triangle := range node.Triangles {
				t, intersection := triangle.GetIntersection(r)
				if t > 0 {
					tValues = append(tValues, t)
					intersections = append(intersections, intersection)
				}
			}

			return tValues, intersections
		}
	}
	return make([]float64, 0), make([][3]float64, 0)
}
