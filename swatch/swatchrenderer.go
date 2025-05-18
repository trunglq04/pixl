package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject // object that we need to draw
	parent  *Swatch             // to make sure whether smth is selected, and implement refreshed()
}

func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

func (renderer SwatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size) // change the background layer for drawing
}

func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square

	}
	canvas.Refresh(renderer.parent) // apply  properties changes
}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *SwatchRenderer) Destroy() {}
