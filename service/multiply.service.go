package service

import (
	"image"
	"image/color"
	"image_processing/global"
	"image_processing/utils"
	"log/slog"
	"runtime"
	"sync"
)

func MultiplyImage(value float64) *image.Image {
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
					pixel := pixels[x][y]
					originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
					if !ok {
						slog.Error("type conversion went wrong")
					}
					newR := uint8(min((float64(originalColor.R) * value), 255))
					newG := uint8(min((float64(originalColor.G) * value), 255))
					newB := uint8(min((float64(originalColor.B) * value), 255))
					col := color.RGBA{
						newR,
						newG,
						newB,
						originalColor.A,
					}
					newImage[x][y] = col
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	multipliedImage := utils.ConvertPixelsToImage(newImage)

	return &multipliedImage
}
