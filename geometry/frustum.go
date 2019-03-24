package geometry

type Frustum struct {
	Origin              [3]float64 //The location of the "top" of the pyramid from which the frustum is constructed (i.e. The camera location)
	Dir                 [3]float64 //The normalized direction vector of the frustum
	Up                  [3]float64 //A normalized vector representing which direction is up for the frustum
	Right               [3]float64 //A normalized vector representing which direction is right for the frustum
	NearPlaneDistance   float64    //Distance to the near plane from Origin
	FarPlaneDistnace    float64    //Distance to the far plane from Origin
	NearPlaneDimensions [2]float64 //Dimensions of the near plane
	FarPlaneDimensions  [2]float64 //Dimensions of the far plane
}

func NewFrustum(origin [3]float64, dir [3]float64, nearPlaneDistance float64, nearPlaneWidth float64, nearPlaneHeight float64) *Frustum {
	res := &Frustum{
		Origin:              origin,
		Dir:                 Normalize(dir),
		NearPlaneDistance:   nearPlaneDistance,
		NearPlaneDimensions: [2]float64{nearPlaneWidth, nearPlaneHeight},
	}

	res.Right = Normalize(Cross(dir, GLOBAL_UP))
	res.Up = Normalize(Cross(res.Right, dir))

	return res
}

func (f *Frustum) Rotate(theta [3]float64) {
	//TODO: Implement this
}

func (f *Frustum) RotateAbout(theta [3]float64, pos [3]float64) {
	//TODO: Implement this
}

func (f *Frustum) LookAt(pos [3]float64) {
	//TODO: Implement this
}

func (f *Frustum) GetRay(x float64, y float64) *Ray {
	//TODO: Implement this
	res := &Ray{
		Origin: f.Origin,
	}

	forward := Scale(f.Dir, f.NearPlaneDistance)
	left := Scale(f.Right, x-(f.NearPlaneDimensions[0]/2))
	up := Scale(f.Up, y-(f.NearPlaneDimensions[1]/2))

	res.Dir = Add(forward, left, up)

	return res
}
