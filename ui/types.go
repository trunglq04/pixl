package ui

import (
	"fyne.io/fyne/v2"
	"trunglq04/pixl/apptype"
	"trunglq04/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
