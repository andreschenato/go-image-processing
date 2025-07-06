package view

import (
	"image_processing/view/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var Window fyne.Window

func MainView(a fyne.App) fyne.Window {
	Window = a.NewWindow("Image Viewer")

	components := container.NewVBox(
		components.ImageComboView(),
		components.Sections(),
	)

	scroll := container.NewCenter(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(components.MinSize().Width+100, components.MinSize().Height+100)),
			container.NewVScroll(
				components,
			),
		),
	)

	content := container.NewBorder(
		scroll,
		nil,
		nil,
		nil,
	)
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(content.MinSize().Width+100, content.MinSize().Height+100))
	return Window
}
