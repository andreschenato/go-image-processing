package view

import (
	"image"
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ImageUploadView() (*fyne.Container, *image.Image) {
	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(250, 250))

	uploadBtn := widget.NewButton("Upload Image", func() {
		service.UploadImage(Window, img)
	})

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(250, 50))
	fixedBtn := container.New(fixedLayout, uploadBtn)

	imgCombo := container.NewVBox(img, fixedBtn)
	return container.NewCenter(imgCombo), &img.Image
}