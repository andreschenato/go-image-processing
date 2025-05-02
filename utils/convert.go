package utils

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

func ConvertImageToPixels(img image.Image) [][]color.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixels := make([][]color.RGBA, width)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, height)
	}

	numCPU := runtime.NumCPU()
	chunks := int(float64(numCPU) * 1.5)

	chunkWidth := (width + chunks - 1) / chunks
	chunkHeight := (height + chunks - 1) / chunks

	var wg sync.WaitGroup

	for startX := 0; startX < width; startX += chunkWidth {
		endX := min(startX + chunkWidth, width)

		for startY := 0; startY < height; startY += chunkHeight {
			endY := min(startY + chunkHeight, height)

			wg.Add(1)
			go func(startX, endX, startY, endY int) {
				defer wg.Done()

				for x := startX; x < endX; x++ {
					for y := startY; y < endY; y++ {
						pixels[x][y] = color.RGBAModel.Convert(img.At(bounds.Min.X+x, bounds.Min.Y+y)).(color.RGBA)
					}
				}
			}(startX, endX, startY, endY)
		}
	}

	wg.Wait()
	return pixels
}

func ConvertPixelsToImage(pixels [][]color.RGBA) image.Image {
	width := len(pixels)
	if width == 0 {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	height := len(pixels[0])
	rect := image.Rect(0, 0, width, height)
	nImg := image.NewRGBA(rect)

	numCPU := runtime.NumCPU()
	chunks := int(float64(numCPU) * 1.5)

	chunkWidth := (width + chunks - 1) / chunks
	chunkHeight := (height + chunks - 1) / chunks

	var wg sync.WaitGroup

	for startX := 0; startX < width; startX += chunkWidth {
		endX := min(startX + chunkWidth, width)

		for startY := 0; startY < height; startY += chunkHeight {
			endY := min(startY + chunkHeight, height)

			wg.Add(1)
			go func(startX, endX, startY, endY int) {
				defer wg.Done()

				for x := startX; x < endX; x++ {
					for y := startY; y < endY; y++ {
						nImg.Set(x, y, pixels[x][y])
					}
				}
			}(startX, endX, startY, endY)
		}
	}

	wg.Wait()
	return nImg
}
