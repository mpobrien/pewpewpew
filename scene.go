package pewpewpew

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Scene struct {
	*image.RGBA
}

func NewScene(width, height int) Scene {
	return Scene{image.NewRGBA(image.Rect(0, 0, width, height))}
}

func (s *Scene) PutPixel(x, y int, rgb Vector) {
	s.Set(x, y, color.RGBA{uint8(rgb.X), uint8(rgb.Y), uint8(rgb.Z), 255})
}

func (s *Scene) Save(filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	return png.Encode(w, s)
}
