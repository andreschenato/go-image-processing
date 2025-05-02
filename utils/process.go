package utils

import (
	"image/color"
	"image_processing/global"
	"log/slog"
)

type PixelTransformFunc func(color.RGBA) color.RGBA
type PixelsTransformFunc func(color.RGBA, color.RGBA) color.RGBA
type AxisTransformFunc func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA)

func single(fun PixelTransformFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	for x := range width {
		for y := range height {
			image[x][y] = fun(pixel[x][y])
		}
	}
	return image
}

func both(fun PixelsTransformFunc, width, height int, pixelOne, pixelTwo, image [][]color.RGBA) [][]color.RGBA {
	for x := range width {
		for y := range height {
			image[x][y] = fun(pixelOne[x][y], pixelTwo[x][y])
		}
	}
	return image
}

func axis(fun AxisTransformFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
	for x := range width {
		for y := range height {
			newX, newY, newColor := fun(x, y, width, height, pixels[x][y])

			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				image[newX][newY] = newColor
			}
		}
	}
	return image
}

func Process(service interface{}) func() {
	return func() {
		if *global.ImageOne == nil && *global.ImageTwo == nil {
			slog.Error("No image set")
			return
		}

		var pixelsTwo *[][]color.RGBA
		imgOne := *global.ImageOne
		if imgOne == nil {
			imgOne = *global.ImageTwo
		}

		pixelsOne := ConvertImageToPixels(imgOne)
		if *global.ImageTwo != nil {
			p := ConvertImageToPixels(*global.ImageTwo)
			pixelsTwo = &p
		}

		xLen := len(pixelsOne)
		yLen := len(pixelsOne[0])

		newImage := make([][]color.RGBA, xLen)
		for i := range xLen {
			newImage[i] = make([]color.RGBA, yLen)
		}

		switch s := service.(type) {
		case PixelTransformFunc:
			newImage = single(s, xLen, yLen, pixelsOne, newImage)
		case PixelsTransformFunc:
			if pixelsTwo != nil {
				newImage = both(s, xLen, yLen, pixelsOne, *pixelsTwo, newImage)
			}
		case AxisTransformFunc:
			newImage = axis(s, xLen, yLen, pixelsOne, newImage)
		default:
			slog.Error("invalid service", "type", s)
			return
		}

		global.FinalImage.Image = ConvertPixelsToImage(newImage)
		global.FinalImage.Refresh()
	}
}
