package utils

import (
	"image/color"
	"image_processing/global"
	"sync"
)

type LowPassFilterFunc func([][]color.RGBA) color.RGBA

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

func mask(pixels [][]color.RGBA, width, height, x, y int) [][]color.RGBA {
	maskSize := global.MaskSize

	neighbors := make([][]color.RGBA, maskSize)
	var wg sync.WaitGroup

	for i := range maskSize {
		neighbors[i] = make([]color.RGBA, maskSize)
	}

	for i := range maskSize {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range maskSize {
				newX := x + i - 1
				newY := y + j - 1

				newX = max(newX, 0)
				newY = max(newY, 0)

				if newX >= width {
					newX = width - 1
				}
				if newY >= height {
					newY = height - 1
				}

				neighbors[i][j] = pixels[newX][newY]
			}
		}(i)
	}

	wg.Wait()
	return neighbors
}
