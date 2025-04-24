package view

import (
	"image_processing/view/buttons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ReverseOperations() *fyne.Container{
	buttons := container.NewCenter(
		container.New(
			layout.NewHBoxLayout(),
			buttons.UpsidedownButton(),
			buttons.MirrorButton(),
		),
	)

	components := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(buttons),
	)

	content := container.NewBorder(container.NewVBox(components), nil, nil, nil)

	return content
}