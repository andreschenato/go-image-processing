package view

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

	content := container.NewGridWithColumns(2, imgOneCont, imgTwoCont)
	return content
}