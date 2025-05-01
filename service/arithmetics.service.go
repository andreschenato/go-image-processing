package service

import (
	"image/color"
	"image_processing/utils"
)

func Sum() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(min(float64(pixelOne.R)+float64(pixelTwo.R), 255)),
			G: uint8(min(float64(pixelOne.G)+float64(pixelTwo.G), 255)),
			B: uint8(min(float64(pixelOne.B)+float64(pixelTwo.B), 255)),
			A: uint8(min(float64(pixelOne.A)+float64(pixelTwo.A), 255)),
		}
	}
}

func Subtract() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(max(float64(pixelOne.R)-float64(pixelTwo.R), 0)),
			G: uint8(max(float64(pixelOne.G)-float64(pixelTwo.G), 0)),
			B: uint8(max(float64(pixelOne.B)-float64(pixelTwo.B), 0)),
			A: pixelOne.A,
		}
	}
}

func Multiply(value float64) utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(min((float64(pixel.R) * value), 255)),
			G: uint8(min((float64(pixel.G) * value), 255)),
			B: uint8(min((float64(pixel.B) * value), 255)),
			A: 255,
		}
	}
}

func Divide(value float64) utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(min((float64(pixel.R) / value), 255)),
			G: uint8(min((float64(pixel.G) / value), 255)),
			B: uint8(min((float64(pixel.B) / value), 255)),
			A: 255,
		}
	}
}

func Diff() utils.PixelsTransformFunc {
	return func(pixelOne, pixelTwo color.RGBA) color.RGBA {
		subOneR := uint8(max(float64(pixelOne.R)-float64(pixelTwo.R), 0))
		subOneG := uint8(max(float64(pixelOne.G)-float64(pixelTwo.G), 0))
		subOneB := uint8(max(float64(pixelOne.B)-float64(pixelTwo.B), 0))

		subTwoR := uint8(max(float64(pixelTwo.R)-float64(pixelOne.R), 0))
		subTwoG := uint8(max(float64(pixelTwo.G)-float64(pixelOne.G), 0))
		subTwoB := uint8(max(float64(pixelTwo.B)-float64(pixelOne.B), 0))

		return color.RGBA{
			R: uint8(min(subOneR+subTwoR, 255)),
			G: uint8(min(subOneG+subTwoG, 255)),
			B: uint8(min(subOneB+subTwoB, 255)),
			A: 255,
		}
	}
}

func Bright(value float64) utils.PixelTransformFunc {
	return func(pixel color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8(max(min(float64(pixel.R)+float64(value), 255), 0)),
			G: uint8(max(min(float64(pixel.G)+float64(value), 255), 0)),
			B: uint8(max(min(float64(pixel.B)+float64(value), 255), 0)),
			A: 255,
		}
	}
}
