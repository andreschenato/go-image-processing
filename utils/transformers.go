package utils

import "image/color"

type PixelTransformFunc func(color.RGBA) color.RGBA
type PixelsTransformFunc func(color.RGBA, color.RGBA) color.RGBA
type AxisTransformFunc func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA)

func single(fun PixelTransformFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixel[x][y])
			}
		}
	})
	return image
}

func both(fun PixelsTransformFunc, width, height int, pixelOne, pixelTwo, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixelOne[x][y], pixelTwo[x][y])
			}
		}
	})
	return image
}

func axis(fun AxisTransformFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				newX, newY, newColor := fun(x, y, width, height, pixels[x][y])

				if newX >= 0 && newX < width && newY >= 0 && newY < height {
					image[newX][newY] = newColor
				}
			}
		}
	})
	return image
}
