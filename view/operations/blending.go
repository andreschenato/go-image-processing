package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Blending() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Average", service.Average()),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			generics.Slider(0, 1, 1, 0.01, "Blending", service.Blend),
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
