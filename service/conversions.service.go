package service

import (
	"image/color"
	"image_processing/utils"
)

func Grayscale() utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		grey := uint8((float64(pixel.R) + float64(pixel.G) + float64(pixel.B)) / 3)
		return color.RGBA{
			R: grey,
			G: grey,
			B: grey,
			A: 255,
		}
	}
}

func Equalize() utils.HistEqualizationFunc {
	return func(pixel color.RGBA) color.RGBA {
		return pixel
	}
}

func Thresholding(threshold float64) utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		grey := uint8((float64(pixel.R) + float64(pixel.G) + float64(pixel.B)) / 3)
		binary := 0
		if grey >= uint8(threshold) {
			binary = 255
		}

		return color.RGBA{
			R: uint8(binary),
			G: uint8(binary),
			B: uint8(binary),
			A: 255,
		}
	}
}
