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
			generics.Button("NOT", service.Not()),
			generics.Button("AND", service.And()),
			generics.Button("OR", service.Or()),
			generics.Button("XOR", service.Xor()),
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
