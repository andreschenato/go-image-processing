package view

import (
	"image/color"
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ProcessedImage() (*fyne.Container) {
	placeholder := canvas.NewRectangle(color.RGBA{40,40,42,255})
	placeholder.SetMinSize(fyne.NewSize(250,250))

	global.FinalImage = canvas.NewImageFromImage(nil)
	global.FinalImage.FillMode = canvas.ImageFillContain
	global.FinalImage.SetMinSize(fyne.NewSize(250, 250))

	return container.NewCenter(container.NewStack(placeholder,global.FinalImage))
}