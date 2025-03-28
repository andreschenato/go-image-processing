package view

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ProcessedImage() (*fyne.Container) {
	global.FinalImage = canvas.NewImageFromImage(nil)
	global.FinalImage.FillMode = canvas.ImageFillContain
	global.FinalImage.SetMinSize(fyne.NewSize(250, 250))

	return container.NewCenter(global.FinalImage)
}