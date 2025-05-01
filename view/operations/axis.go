package operations

import (
	"image_processing/service"
	"image_processing/view/generics"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Axis() *fyne.Container{
	buttons := container.NewCenter(
		container.NewHBox(
			generics.Button("Vertical", service.Vertical()),
			generics.Button("Horizontal", service.Horizontal()),
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