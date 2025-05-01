package view

import (
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func LogicOperations() *fyne.Container {
	buttons := container.NewCenter(
		container.NewHBox(
			Button("Not", service.Not()),
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
