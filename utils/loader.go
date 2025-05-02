package utils

import (
	"errors"
	"image"
	_ "image/png"
	_ "image/gif"
	"log/slog"
	"os"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("error opening file", "error", err)
		return nil, err
	}

	img, format, err := image.Decode(f)
	if err != nil {
		slog.Error("Decoding error:", "error", err.Error())
		return nil, err
	}

	if format != "jpeg" && format != "png" && format != "gif" && format != "tiff" && format != "bmp" {
		slog.Error("image format is not valid")
		return nil, errors.New("image format is not valid")
	}
	return img, nil
}
