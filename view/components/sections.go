package components

import (
	"image_processing/view/operations"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Sections() *fyne.Container {
	opts := []string{"Arithmetics", "Conversions", "Axis", "Blending", "Logic", "Low-Pass", "High-Pass"}
	var components fyne.Container
	selection := widget.NewSelect(opts[:], func(s string) {
		switch s {
		case "Arithmetics":
			components = *operations.Arithmetics()
		case "Conversions":
			components = *operations.Conversions()
		case "Axis":
			components = *operations.Axis()
		case "Blending":
			components = *operations.Blending()
		case "Logic":
			components = *operations.Logics()
		case "Low-Pass":
			components = *operations.LowPassFilters()
		case "High-Pass":
			components = *operations.HighPassFilters()
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
			container.NewCenter(selectionLayout),
			container.NewVBox(&components),
		),
	)

	return container.NewPadded(contents)
}
