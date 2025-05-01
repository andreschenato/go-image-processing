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

func And() utils.PixelsTransformFunc {
	return func(pixelOne color.RGBA, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(pixelOne.R & pixelTwo.R),
			G: uint8(pixelOne.G & pixelTwo.G),
			B: uint8(pixelOne.B & pixelTwo.B),
			A: 255,
		}
	}
}

func Or() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(pixelOne.R | pixelTwo.R),
			G: uint8(pixelOne.G | pixelTwo.G),
			B: uint8(pixelOne.B | pixelTwo.B),
			A: 255,
		}
	}
}

func Xor() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8((pixelOne.R | pixelTwo.R) & 255 - (pixelOne.R & pixelTwo.R)),
			G: uint8((pixelOne.G | pixelTwo.G) & 255 - (pixelOne.G & pixelTwo.G)),
			B: uint8((pixelOne.B | pixelTwo.B) & 255 - (pixelOne.B & pixelTwo.B)),
			A: 255,
		}
	}
}
