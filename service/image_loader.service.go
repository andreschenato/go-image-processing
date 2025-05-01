package service

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	_ "image/gif"
	"log/slog"
	"os"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println(fi.Name())

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
