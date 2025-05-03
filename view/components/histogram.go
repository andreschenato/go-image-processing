package components

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Histogram() *fyne.Container {
	global.Hist = canvas.NewImageFromImage(nil)
	global.Hist.FillMode = canvas.ImageFillContain
	global.Hist.SetMinSize(fyne.NewSize(800, 250))

	return container.NewCenter(global.Hist)
}
