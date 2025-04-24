package view

import (
	"image_processing/view/sliders"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func BlendingOperations() *fyne.Container{
	buttons := container.NewCenter(
		container.New(
			layout.NewHBoxLayout(),
			sliders.BlendingRatioSlider(),
		),
	)

	components := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(buttons),
	)

	content := container.NewBorder(container.NewVBox(components), nil, nil, nil)

	return content
}