package renderer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/dkkempto/freed/geometry"
)

type Renderer struct {
	Camera          *Camera
	PathToInput     string
	PathToOutput    string
	NumberOfCameras int
	Resolution      [2]int
	ThetaStep       float64
}

func NewRenderer(pathToInputFile string, pathToOutputFile string, numCameras int, numPixelsX int, numPixelsY int, dTheta float64) *Renderer {
	//TODO: Return a renderer :)

	return nil
}

func (r *Renderer) Render(mesh *geometry.Mesh) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{r.Camera.ScreenResolution[0], r.Camera.ScreenResolution[1]}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < r.Camera.ScreenResolution[0]; x++ {
		for y := 0; y < r.Camera.ScreenResolution[1]; y++ {
			tValues, _ := mesh.GetIntersections(r.Camera.GetRay(x, y))

			intensity := 0.0
			prevIntensity := 0.0
			prevT := 0.0

			if len(tValues) > 0 {
				for i, t := range tValues {
					if i%2 == 0 {
						// intensity = prevIntensity + (1.0 - t)
						prevT = t
					} else {
						intensity = prevIntensity + (t - prevT)
						prevIntensity = intensity
						prevT = t
					}
				}
				//																										: [sum=0.0 prevSum=0.0]
				//[i=0 t=0.2 prevT=0.0 prevSum=0.0] 0.0 + (1.0 - 0.2) : [sum=0.8 prevSum=0.0]
				//[i=1 t=0.4 prevT=0.2 prevSum=0.0] 0.0 + (0.4 - 0.2) : [sum=0.2 prevSum=0.2]
				//[i=2 t=0.6 prevT=0.4 prevSum=0.2] 0.2 + (1.0 - 0.6) : [sum=0.6 prevSum=0.2]
				//[i=3 t=0.8 prevT=0.6 prevSum=0.2] 0.2 + (0.8 - 0.6) : [sum=0.4 prevSum=0.4]
				i := uint8(intensity * 255)

				img.SetRGBA(x, y, color.RGBA{i, i, i, 0xff})
			}
		}
	}

	fileName := fmt.Sprintf("%v/test_%04d.png", r.PathToOutput, 0)
	f, _ := os.Create(fileName)
	png.Encode(f, img)
	f.Close()
	//TODO: Render the mesh and output to folder
}
