package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Window fyne.Window

func MainView() fyne.Window {
	myApp := app.NewWithID("Process image")
	Window = myApp.NewWindow("Image Viewer")

	content := ImageComboView()
	Window.SetContent(content)

	Window.Resize(fyne.NewSize(1000, 400))
	return Window
}