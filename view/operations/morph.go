package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Morph() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Dilation", service.Dilation()),
			generics.Button("Erosion", service.Erosion()),
			generics.Button("Opening", service.Opening()),
			generics.Button("Closing", service.Closing()),
			generics.Button("Outline", service.Outline()),
		),
	)

	radios := container.NewCenter(
		container.NewHBox(
			generics.MaskFormat(),
		),
	)

	sliders := container.NewCenter(
		container.NewVBox(
			generics.MaskSlider(),
		),
	)

	components := container.NewCenter(
		container.NewVBox(
			buttons,
			radios,
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
