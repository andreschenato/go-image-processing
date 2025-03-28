package view

import (
	"image_processing/view/buttons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var Window fyne.Window

func MainView(a fyne.App) fyne.Window {
	Window = a.NewWindow("Image Viewer")

	components := container.NewHBox(ImageComboView(), buttons.GrayscaleButton(), ProcessedImage())
	content := container.NewCenter(components)
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(1000, 400))
	return Window
}