package sdlEngine

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func Draw(screen *sdl.Surface, pngImagePath string, x int32, y int32) (err error) {
	// Load a PNG image
	var pngImage *sdl.Surface
	if pngImage, err = img.Load(pngImagePath); err != nil {
		return err
	}
	defer pngImage.Free()

	// Draw the PNG image on the first half of the window
	pngImage.BlitScaled(nil, screen, &sdl.Rect{X: x, Y: y, W: 16, H: 16}) //TODO: separate draw coordinate and game coordinate

	return
}
