package service

import (
	"image/color"
	"image_processing/utils"
)

func Not() utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(255 - pixel.R),
			G: uint8(255 - pixel.G),
			B: uint8(255 - pixel.B),
			A: 255,
		}
	}
}