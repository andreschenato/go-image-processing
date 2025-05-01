package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Arithmetics() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Sum", service.Sum()),
			generics.Button("Subtract", service.Subtract()),
			generics.Button("Diff", service.Diff()),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			generics.Slider(-100, 100, 0, 1, "Bright", service.Bright),
			generics.Slider(0, 2, 1, 0.01, "Multiply", service.Multiply),
			generics.Slider(0, 8, 1, 0.01, "Divide", service.Divide),
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
