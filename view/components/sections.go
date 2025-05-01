package components

import (
	"image_processing/service"
	"image_processing/view/generics"
	"image_processing/view/operations"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Sections() *fyne.Container {
	opts := []string{"Arithmetics", "Grayscale", "Axis", "Blending", "Logic"}
	var components fyne.Container
	selection := widget.NewSelect(opts[:], func(s string) {
		switch s {
		case "Arithmetics":
			components = *operations.Arithmetics()
		case "Grayscale":
			components = *generics.Button("Grayscale", service.Grayscale())
		case "Axis":
			components = *operations.Axis()
		case "Blending":
			components = *operations.Blending()
		case "Logic":
			components = *operations.Logics()
		}
	})

	selectionLayout := container.New(
		layout.NewGridWrapLayout(
			fyne.NewSize(
				200,
				50,
			),
		),
		selection,
	)

	contents := container.NewCenter(
		container.NewVBox(
			generics.Checkbox("Process singlethread?"),
			container.NewCenter(selectionLayout),
			container.NewVBox(&components),
		),
	)

	return container.NewPadded(contents)
}
