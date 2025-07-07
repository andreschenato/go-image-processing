package utils

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func warning(label string) {
	closeButton := widget.NewButton("Close", nil)

	content := widget.NewCard("Warning", label,
		container.NewCenter(
			container.NewHBox(closeButton),
		),
	)

	popup := widget.NewModalPopUp(content, global.Window.Canvas())
	closeButton.OnTapped = func() {
		popup.Hide()
	}

	popup.Resize(fyne.NewSize(250, 50))

	popup.Show()
}
