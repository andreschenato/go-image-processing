package utils

import (
	"image/color"
)

type LowPassFilterFunc func([][]color.RGBA) color.RGBA
type HighPassFilterFunc func([][]color.RGBA) color.RGBA

func lowPass(fun LowPassFilterFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				maskPixels := mask(pixel, width, height, x, y)
				image[x][y] = fun(maskPixels)
			}
		}
	})
	return image
}

func highPass(fun HighPassFilterFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				maskPixels := mask(pixel, width, height, x, y)
				image[x][y] = fun(maskPixels)
			}
		}
	})
	return image
}
