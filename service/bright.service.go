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

func BrightImage(brightValue float64) image.Image {
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

			for x := range endX {
				for y := range yLen {
					pixel := pixels[x][y]
					originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
					if !ok {
						slog.Error("type conversion went wrong")
					}

					newR := uint8(max(min(float64(originalColor.R)+float64(brightValue), 255), 0))
					newG := uint8(max(min(float64(originalColor.G)+float64(brightValue), 255), 0))
					newB := uint8(max(min(float64(originalColor.B)+float64(brightValue), 255), 0))

					newImage[x][y] = color.RGBA{
						newR,
						newG,
						newB,
						originalColor.A,
					}

				}
			}
		}(startX, endX)
	}
	wg.Wait()

	brightenImage := utils.ConvertPixelsToImage(newImage)

	return brightenImage
}
