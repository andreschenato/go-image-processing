package generics

import (
	"image_processing/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Button(label string, service interface{}) *fyne.Container {
	btn := widget.NewButton(label, utils.Process(service))

	btn.Resize(fyne.NewSize(btn.MinSize().Width+10, btn.MinSize().Height+5))

	return container.NewPadded(
		container.NewCenter(btn),
	)
}
