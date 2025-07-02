package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func LowPassFilters() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Min", service.Min()),
			generics.Button("Max", service.Max()),
			generics.Button("Mean", service.Mean()),
			generics.Button("Median", service.Median()),
			generics.Button("Conservative", service.Conservative()),
		),
	)

	sliders := generics.FilterSlider()

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
