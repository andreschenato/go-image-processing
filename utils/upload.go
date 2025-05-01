package utils

import (
	"log/slog"

	"fyne.io/fyne/v2/canvas"
	"github.com/sqweek/dialog"
)

func UploadImage(image *canvas.Image) {
	file, err := dialog.File().Filter("Image file", "jpeg", "jpg", "png", "gif", "tif", "bmp").Load()
	if err != nil {
		slog.Error("Fatal error opening file dialog")
		return
	}

	image.Image, err = LoadImage(file)
	if err != nil {
		slog.Error("Fatal error loading image")
	}
	image.Refresh()
}
