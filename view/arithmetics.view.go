package view

import (
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ArithmeticOperations() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			Button("Sum", service.Sum()),
			Button("Subtract", service.Subtract()),
			Button("Diff", service.Diff()),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			Slider(-100, 100, 0, 1, "Bright", service.Bright),
			Slider(0, 2, 1, 0.01, "Multiply", service.Multiply),
			Slider(0, 8, 1, 0.01, "Divide", service.Divide),
		),
	)

	components := container.NewCenter(
		container.NewVBox(
			buttons,
			sliders,
		),
	)

	content := container.NewBorder(
		components,
		nil,
		nil,
		nil,
	)

	return content
}
