package service

import (
	"log/slog"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
)

func UploadImage(w fyne.Window, image *canvas.Image) {
	dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			slog.Error("Fatal error opening file dialog")
			return
		}
		if f == nil {
			return
		}

		imgPath := f.URI().Path()

		image.Image, err = LoadImage(imgPath)
		if err != nil {
			slog.Error("Fatal error loading image")
			dialog.NewError(err, w).Resize(fyne.NewSize(200, 50))
		}
		image.Refresh()
	}, w)
}
