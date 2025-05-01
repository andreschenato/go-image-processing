package generics

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Checkbox(label string) *fyne.Container{
	check := widget.NewCheck(label, func(b bool) { global.UseSingleThread = b})

	check.Resize(fyne.NewSize(check.MinSize().Width+10, check.MinSize().Height+5))

	return container.NewPadded(
		container.NewCenter(check),
	)
} 