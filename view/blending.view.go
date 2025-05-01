package view

import (
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BlendingOperations() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			Button("Average", service.Average()),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			Slider(0, 1, 1, 0.01, "Blending", service.Blend),
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
