package service

import (
	"image/color"
	"image_processing/utils"
)

func Blend(ratio float64) utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(min(ratio*float64(pixelOne.R)+(1-ratio)*float64(pixelTwo.R), 255)),
			G: uint8(min(ratio*float64(pixelOne.G)+(1-ratio)*float64(pixelTwo.G), 255)),
			B: uint8(min(ratio*float64(pixelOne.B)+(1-ratio)*float64(pixelTwo.B), 255)),
			A: 255,
		}
	}
}

func Average() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(min(min(float64(pixelOne.R)+float64(pixelTwo.R), 255)/2, 255)),
			G: uint8(min(min(float64(pixelOne.G)+float64(pixelTwo.G), 255)/2, 255)),
			B: uint8(min(min(float64(pixelOne.B)+float64(pixelTwo.B), 255)/2, 255)),
			A: 255,
		}
	}
}
