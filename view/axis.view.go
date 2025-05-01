package view

import (
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func AxisOperations() *fyne.Container{
	buttons := container.NewCenter(
		container.NewHBox(
			Button("Vertical", service.Vertical()),
			Button("Horizontal", service.Horizontal()),
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