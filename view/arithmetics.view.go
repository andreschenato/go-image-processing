package view

import (
	"image_processing/view/buttons"
	"image_processing/view/sliders"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ArithmeticOperations() *fyne.Container {
	buttons := container.NewCenter(
		container.New(
			layout.NewHBoxLayout(),
			buttons.SumButton(),
			buttons.SubtractButton(),
			buttons.DiffButton(),
		),
	)

	components := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(buttons),
		container.NewPadded(sliders.BrightSlider()),
		container.NewPadded(sliders.MultiplySlider()),
		container.NewPadded(sliders.DivideSlider()),
	)

	content := container.NewBorder(container.NewVBox(components), nil, nil, nil)

	return content
}
