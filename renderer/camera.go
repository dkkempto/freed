package renderer

import "github.com/dkkempto/freed/geometry"

type Camera struct {
	Origin           [3]float64 //Physical location of the camera with respect to the origin about which it rotates
	Dir              [3]float64 //The normalized direction the camera is facing
	Up               [3]float64 //The local up vector
	Right            [3]float64 //The local right vector
	Forward          [3]float64 //The direction vector scaled to screen size
	ThrowDistance    float64    //Distance from the camera to the virtual screen to which it is projecting
	ScreenDimensions [2]float64 //The physical dimensions of the screen
	ScreenResolution [2]int     //The pixel resolution of the projector used for printing
}

//Sort of a nasty constructor :( there's probs a better way to do this?
func NewCamera(origin [3]float64, dir [3]float64, throwDistance float64, screenDimensions [2]float64, screenResolution [2]int) *Camera {
	res := &Camera{
		Origin:           origin,
		Dir:              geometry.Normalize(dir),
		ThrowDistance:    throwDistance,
		ScreenDimensions: screenDimensions,
		ScreenResolution: screenResolution,
	}
	res.Forward = geometry.Scale(res.Dir, throwDistance)
	res.Right = geometry.Normalize(geometry.Cross(res.Dir, geometry.GLOBAL_UP))
	res.Up = geometry.Normalize(geometry.Cross(res.Right, res.Dir))

	return res
}

func (c *Camera) GetRay(x int, y int) *geometry.Ray {
	/*
		The direction vector of the returned ray should not be normalized. By looking at the t-values of the
		intersections, we can adjust the brightness accordingly. Say, t = 0.2, t = 0.4, we can take the range
		0.2 and use that as our brightness value. Honestly could probably scale this accordingly as well,
		by taking the range of possible intersections based on the tank being printed in, because we would
		never get a full intensity otherwise.
	*/

	u := (float64(x) - float64(c.ScreenResolution[0])/2) / float64(c.ScreenResolution[0])

	//(0, 0) is at top of screen, so if y = 0, we should move up instead of down
	v := (float64(c.ScreenResolution[1])/2 - float64(y)) / float64(c.ScreenResolution[1])

	left := geometry.Scale(c.Right, u*c.ScreenDimensions[0])
	up := geometry.Scale(c.Up, v*c.ScreenDimensions[1])

	return &geometry.Ray{
		Origin: c.Origin,
		Dir:    geometry.Add(c.Forward, left, up),
	}
}
