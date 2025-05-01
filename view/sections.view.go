package view

import (
	"fmt"
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Sections() *fyne.Container {
	opts := [5]string{"Arithmetics", "Grayscale", "Reverse", "Blending", "Logic"}
	var components fyne.Container
	selection := widget.NewSelect(opts[:], func(s string) {
		fmt.Println(s)
		switch s {
		case "Arithmetics":
			components = *ArithmeticOperations()
		case "Grayscale":
			components = *Button("Grayscale", service.Grayscale())
		case "Axis":
			components = *AxisOperations()
		case "Blending":
			components = *BlendingOperations()
		case "Logic":
			components = *LogicOperations()
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

	contents := container.New(
		layout.NewVBoxLayout(),
		container.NewCenter(selectionLayout),
		container.NewVBox(&components),
	)

	return container.NewCenter(contents)
}
