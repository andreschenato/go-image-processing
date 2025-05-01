package components

import (
	"image/color"
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ImageComboView() *fyne.Container {
	imgOneCont, imgOne := ImageUploadView()
	imgTwoCont, imgTwo := ImageUploadView()

	if imgOne != nil {
		global.ImageOne = imgOne
	}
	if imgTwo != nil {
		global.ImageTwo = imgTwo
	}

	var equals = canvas.NewText("=", color.White)
	var imgA = canvas.NewText("Imagem A", color.White)
	var imgB = canvas.NewText("Imagem B", color.White)
	var final = canvas.NewText("Resultado", color.White)

	equals.TextSize = 48
	imgA.TextSize = 36
	imgB.TextSize = 36
	final.TextSize = 36

	content := container.New(
		layout.NewHBoxLayout(),
		container.NewVBox(
			container.NewCenter(imgA),
			container.NewPadded(imgOneCont),
		),
		container.NewVBox(
			container.NewCenter(imgB),
			container.NewPadded(imgTwoCont),
		),
		container.NewCenter(
			container.NewPadded(equals),
		),
		container.NewVBox(
			container.NewCenter(final),
			container.NewPadded(ProcessedImage()),
		),
	)
	return container.NewCenter(content)
}
