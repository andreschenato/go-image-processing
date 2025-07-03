package utils

import (
	"image/color"
)

type MorphOpsFunc func([][]color.RGBA) color.RGBA

func morph(fun MorphOpsFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
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
