package view

import (
	"image/color"
	"image_processing/global"
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ProcessedImage() (*fyne.Container) {
	placeholder := canvas.NewRectangle(color.RGBA{40,40,42,255})
	placeholder.SetMinSize(fyne.NewSize(250,250))

	global.FinalImage = canvas.NewImageFromImage(nil)
	global.FinalImage.FillMode = canvas.ImageFillContain
	global.FinalImage.SetMinSize(fyne.NewSize(250, 250))

	downloadBtn := widget.NewButton("Download Image", func () {
		service.DownloadImage(Window)
	})

	imgContainer := container.NewStack(placeholder, global.FinalImage)

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(250, 50))
	fixedBtn := container.New(fixedLayout, downloadBtn)

	return container.NewVBox(
		container.NewPadded(imgContainer),
		container.NewPadded(fixedBtn),
	)
}