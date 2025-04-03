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

func GrayscaleImage() *image.Image {
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
					grey := uint8((float64(originalColor.R) + float64(originalColor.G) + float64(originalColor.B)) / 3)
					col := color.RGBA{
						grey,
						grey,
						grey,
						originalColor.A,
					}
					newImage[x][y] = col
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	grayscaledImage := utils.ConvertPixelsToImage(newImage)

	return &grayscaledImage
}
