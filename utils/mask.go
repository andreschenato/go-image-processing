package utils

import (
	"image/color"
	"sync"
)

func mask(pixels [][]color.RGBA, width, height, x, y int, maskSize int) [][]color.RGBA {
	offset := maskSize / 2

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
				newX := x + i - offset
				newY := y + j - offset

				if newX < 0 {
					newX = 0
				}
				if newY < 0 {
					newY = 0
				}
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
