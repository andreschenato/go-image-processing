package service

import (
	"fmt"
	"image"
	"image/color"
	"image_processing/global"
	"image_processing/utils"
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

	wg := sync.WaitGroup{}
	for x := range xLen {
		for y := range yLen {
			wg.Add(1)
			go func(x, y int) {
				pixel := pixels[x][y]
				originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok {
					fmt.Println("type conversion went wrong")
				}
				grey := uint8((float64(originalColor.R) + float64(originalColor.G) + float64(originalColor.B)) / 3)
				col := color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage[x][y] = col
				wg.Done()
			}(x, y)

		}
	}
	wg.Wait()

	grayscaledImage := utils.ConvertPixelsToImage(newImage)

	return &grayscaledImage
}
