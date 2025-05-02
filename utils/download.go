package utils

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"image_processing/global"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

func DownloadImage() {
	if global.FinalImage.Image == nil {
		return
	}

	file, err := dialog.File().Filter("Save image file", "jpeg", "jpg", "png", "gif", "tif", "bmp").
		Title("Export image to file").Save()
	if err != nil {
		slog.Error("Error trying to save file")
		return
	}

	f, err := os.Create(file)
	if err != nil {
		slog.Error("Error trying to create file", "error", err)
	}
	defer f.Close()

	switch filepath.Ext(file) {
	case ".jpeg", ".jpg":
		if err := jpeg.Encode(f, global.FinalImage.Image, nil); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
	case ".png":
		if err := png.Encode(f, global.FinalImage.Image); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
	case ".gif":
		if err := gif.Encode(f, global.FinalImage.Image, nil); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
	case ".tif":
		if err := tiff.Encode(f, global.FinalImage.Image, nil); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
	case ".bmp":
		if err := bmp.Encode(f, global.FinalImage.Image); err != nil {
			slog.Error("Failed to encode image", "error", err)
			return
		}
	}
}
