package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Logics() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Not", service.Not()),
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
