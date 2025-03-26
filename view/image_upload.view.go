package view

import (
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ImageUploadView() *fyne.Container {
	image := canvas.NewImageFromImage(nil)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(250, 250))

	uploadBtn := widget.NewButton("Upload Image", func() {
		service.UploadImage(Window, image)
	})

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(250, 50))
	fixedBtn := container.New(fixedLayout, uploadBtn)

	imgCombo := container.NewVBox(image, fixedBtn)
	return container.NewCenter(imgCombo)
}