package service

import (
	"log/slog"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	fdialog "fyne.io/fyne/v2/dialog"
	"github.com/sqweek/dialog"
)

func UploadImage(w fyne.Window, image *canvas.Image) {
	file, err := dialog.File().Filter("Image file", "jpeg", "jpg", "png", "gif", "tif", "bmp").Load()
	if err != nil {
		slog.Error("Fatal error opening file dialog")
		return
	}

	image.Image, err = LoadImage(file)
	if err != nil {
		slog.Error("Fatal error loading image")
		errDialog := fdialog.NewError(err, w)
		errDialog.Resize(fyne.NewSize(200, 50))
		errDialog.Show()
	}
	image.Refresh()
}
