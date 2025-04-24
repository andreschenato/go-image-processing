package service

import (
	"image"
	"image/color"
	"image_processing/global"
	"image_processing/utils"
	"runtime"
	"sync"
)

func UpsidedownImage() *image.Image {
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := utils.ConvertImageToPixels(img)

	xLen := len(pixels)
	yLen := len(pixels[0])

	newImage := make([][]color.Color, xLen)
	for i := range newImage {
		newImage[i] = make([]color.Color, yLen)
	}

	numWorkers := runtime.NumCPU()
	var wg = sync.WaitGroup{}

	chunkSize := max(xLen/numWorkers, 1)

	for i := range numWorkers {
		wg.Add(1)

		startX := i * chunkSize
		endX := startX + chunkSize
		if i == numWorkers-1 {
			endX = xLen
		}

		go func(startX, endX int) {
			defer wg.Done()

			for x := range xLen {
				for y := range yLen {
					newImage[x][yLen-1-y] = pixels[x][y]
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	upsidedownImage := utils.ConvertPixelsToImage(newImage)

	return &upsidedownImage
}
