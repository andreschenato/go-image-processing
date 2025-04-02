package view

import (
	"image_processing/view/buttons"
	"image_processing/view/sliders"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var Window fyne.Window

func MainView(a fyne.App) fyne.Window {
	Window = a.NewWindow("Image Viewer")

	components := container.New(
		layout.NewVBoxLayout(),
		ImageComboView(),
		container.NewPadded(buttons.GrayscaleButton()),
		container.NewPadded(buttons.SumButton()),
		container.NewPadded(sliders.BrightSlider()),
		ProcessedImage(),
	)
	content := container.NewCenter(components)
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(1000, 400))
	return Window
}