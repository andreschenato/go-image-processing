package view

import (
	"image_processing/view/buttons"
	"image_processing/view/sliders"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func BlendingOperations() *fyne.Container{
	buttons := container.NewCenter(
		container.New(
			layout.NewHBoxLayout(),
			buttons.AverageButton(),
		),
	)

	components := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(buttons),
		container.NewPadded(sliders.BlendingRatioSlider()),
	)

	content := container.NewBorder(container.NewVBox(components), nil, nil, nil)

	return content
}