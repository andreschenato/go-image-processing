package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func HighPassFilters() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Sobel", service.Sobel()),
			generics.Button("Prewitt", service.Prewitt()),
			generics.Button("Laplacian Exterior", service.LaplacianExterior()),
			generics.Button("Laplacian Interior", service.LaplacianInterior()),
		),
	)

	components := container.NewCenter(
		container.NewVBox(
			buttons,
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
