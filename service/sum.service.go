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

func SumImages() *image.Image {
	if global.ImageOne == nil || global.ImageTwo == nil {
		slog.Error("Images aren't set")
	}

	imgOne := *global.ImageOne
	imgTwo := *global.ImageTwo

	pixelsOne := utils.ConvertImageToPixels(imgOne)
	pixelsTwo := utils.ConvertImageToPixels(imgTwo)

	xLen := len(pixelsOne)
	yLen := len(pixelsOne[0])

	newImage := make([][]color.Color, xLen)
	for i := range xLen {
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
					pixelOne := pixelsOne[x][y]
					pixelTwo := pixelsTwo[x][y]

					ogColorOne, ok := color.RGBAModel.Convert(pixelOne).(color.RGBA)
					if !ok {
						slog.Error("conversion went wrong")
					}

					ogColorTwo, ok := color.RGBAModel.Convert(pixelTwo).(color.RGBA)
					if !ok {
						slog.Error("conversion went wrong")
					}

					newR := uint8(min(float64(ogColorOne.R)+float64(ogColorTwo.R), 255))
					newG := uint8(min(float64(ogColorOne.G)+float64(ogColorTwo.G), 255))
					newB := uint8(min(float64(ogColorOne.B)+float64(ogColorTwo.B), 255))
					newA := uint8(min(float64(ogColorOne.A)+float64(ogColorTwo.A), 255))

					newImage[x][y] = color.RGBA{
						newR,
						newG,
						newB,
						newA,
					}
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	resultImage := utils.ConvertPixelsToImage(newImage)

	return &resultImage
}
