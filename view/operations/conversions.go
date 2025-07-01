package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Conversions() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Grayscale", service.Grayscale()),
			generics.Button("Equalize", service.Equalize()),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			generics.Slider(0, 255, 0, 1, "Thresholding", service.Thresholding),
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
