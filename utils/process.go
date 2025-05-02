package utils

import (
	"image/color"
	"image_processing/global"
	"log/slog"
)

type PixelTransformFunc func(color.RGBA) color.RGBA
type PixelsTransformFunc func(color.RGBA, color.RGBA) color.RGBA
type AxisTransformFunc func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA)

func single(fun PixelTransformFunc) {
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

	xLen := len(pixels)
	yLen := len(pixels[0])

	newImage := make([][]color.RGBA, xLen)
	for i := range newImage {
		newImage[i] = make([]color.RGBA, yLen)
	}

	for x := range xLen {
		for y := range yLen {
			pixel := pixels[x][y]

			newColor := fun(pixel)

			newImage[x][y] = newColor
		}
	}

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
}

func both(fun PixelsTransformFunc) {
	pixelsOne := ConvertImageToPixels(*global.ImageOne)
	pixelsTwo := ConvertImageToPixels(*global.ImageTwo)

	xLen := len(pixelsOne)
	yLen := len(pixelsOne[0])

	newImage := make([][]color.RGBA, xLen)
	for i := range xLen {
		newImage[i] = make([]color.RGBA, yLen)
	}

	for x := range xLen {
		for y := range yLen {
			newColor := fun(pixelsOne[x][y], pixelsTwo[x][y])

			newImage[x][y] = newColor
		}
	}

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
}

func axis(fun AxisTransformFunc) {
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

	xLen := len(pixels)
	yLen := len(pixels[0])

	newImage := make([][]color.RGBA, xLen)
	for i := range newImage {
		newImage[i] = make([]color.RGBA, yLen)
	}

	for x := range xLen {
		for y := range yLen {
			pixel := pixels[x][y]

			newX, newY, newColor := fun(x, y, xLen, yLen, pixel)

			if newX >= 0 && newX < xLen && newY >= 0 && newY < yLen {
				newImage[newX][newY] = newColor
			}
		}
	}

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
}

func Process(service interface{}) func() {
	return func() {
		switch s := service.(type) {
		case PixelTransformFunc:
			if *global.ImageOne != nil || *global.ImageTwo != nil {
				single(s)
				return
			}
			slog.Error("No image set")
			return
		case PixelsTransformFunc:
			if *global.ImageOne != nil && *global.ImageTwo != nil {
				both(s)
				return
			}
			slog.Error("Images aren't set")
			return
		case AxisTransformFunc:
			if *global.ImageOne != nil || *global.ImageTwo != nil {
				axis(s)
				return
			}
			slog.Error("No image set")
			return
		default:
			slog.Error("invalid service", "type", s)
			return
		}
	}
}
