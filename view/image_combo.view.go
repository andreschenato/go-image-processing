package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ImageComboView() *fyne.Container {
	imgUploader1 := ImageUploadView()
	imgUploader2 := ImageUploadView()

	content := container.NewGridWithColumns(2, imgUploader1, imgUploader2)
	return content
}