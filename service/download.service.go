package service

import (
	"image/png"
	"image_processing/global"
	"image_processing/view/popups"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DownloadImage(w fyne.Window) {
	if global.FinalImage.Image == nil {
		return
	}

	dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
		if err != nil {
			slog.Error("Fatal error saving file")
			return
		}
		if f == nil {
			return
		}

		path := f.URI().Path()

		if err := png.Encode(f, global.FinalImage.Image); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
		if err := f.Close(); err != nil {
			slog.Error("Failed to close file", "error", err)
			return
		}
		if f.URI().Extension() != ".png" {
			newPath := path + ".png"
			os.Rename(path, newPath)
		}

		popups.SavedPopup(w)
	}, w)
}
