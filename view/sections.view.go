package view

import (
	"fmt"
	"image_processing/view/buttons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Sections() *fyne.Container {
	opts := [2]string{"Arithmetics", "Grayscale"}
	var components fyne.Container
	selection := widget.NewSelect(opts[:], func(s string) {
		fmt.Println(s)
		switch s {
		case "Arithmetics":
			components = *ArithmeticOperations()
		case "Grayscale":
			components = *buttons.GrayscaleButton()
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
