package view

import (
	"image_processing/view/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var Window fyne.Window

func MainView(a fyne.App) fyne.Window {
	Window = a.NewWindow("Image Viewer")

	components := container.NewVBox(
		components.ImageComboView(),
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
