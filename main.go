package main

import (
	"image_processing/view"

	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.NewWithID("process_image_app")
	view.MainView(myApp).ShowAndRun()
}
