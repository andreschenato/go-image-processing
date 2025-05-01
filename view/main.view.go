package view

import (
	"image/color"
	"image_processing/global"
	"image_processing/view/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var Window fyne.Window

func MainView(a fyne.App) fyne.Window {
	Window = a.NewWindow("Image Viewer")

	components := container.NewVBox(
		components.ImageComboView(),
		container.NewCenter(
			container.NewHBox(
				canvas.NewText("Execution time:", color.White),
				&global.ExecutionTime,
			),
		),
		components.Sections(),
	)

	content := container.NewBorder(
		components,
		nil,
		nil,
		nil,
	)
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(content.MinSize().Width+100, content.MinSize().Height+100))
	return Window
}
