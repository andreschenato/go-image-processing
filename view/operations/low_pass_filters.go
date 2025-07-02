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
			generics.Button("MIN", service.Min()),
			generics.Button("MAX", service.Max()),
			generics.Button("MEAN", service.Mean()),
			generics.Button("Median", service.Median()),
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
