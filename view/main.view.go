package view

import (
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
		container.NewPadded(Sections()),
	)

	content := container.NewBorder(container.NewVBox(components), nil, nil, nil)
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(content.MinSize().Width + 100, content.MinSize().Height + 100))
	return Window
}