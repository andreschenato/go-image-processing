package utils

import (
	"image/color"
	"image_processing/global"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

type PixelTransformFunc func(color.RGBA) color.RGBA
type PixelsTransformFunc func(color.RGBA, color.RGBA) color.RGBA
type AxisTransformFunc func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA)

func parallelizeProcessing(width, height int, process func(startX, endX, height int)) {
	start := time.Now()
	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numCPU)

	chunkSize := max(width/numCPU, 1)

	for i := range numCPU {
		startX := i * chunkSize
		endX := (i + 1) * chunkSize
		if i == numCPU-1 {
			endX = width
		}

		go func(startX, endX int) {
			defer wg.Done()
			process(startX, endX, height)
		}(startX, endX)
	}

	wg.Wait()
	global.ExecutionTime.SetText(time.Since(start).String())
}

func single(fun PixelTransformFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	start := time.Now()
	for x := range width {
		for y := range height {
			image[x][y] = fun(pixel[x][y])
		}
	}
	global.ExecutionTime.SetText(time.Since(start).String())
	return image
}

func both(fun PixelsTransformFunc, width, height int, pixelOne, pixelTwo, image [][]color.RGBA) [][]color.RGBA {
	start := time.Now()
	for x := range width {
		for y := range height {
			image[x][y] = fun(pixelOne[x][y], pixelTwo[x][y])
		}
	}
	global.ExecutionTime.SetText(time.Since(start).String())
	return image
}

func axis(fun AxisTransformFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
	start := time.Now()
	for x := range width {
		for y := range height {
			newX, newY, newColor := fun(x, y, width, height, pixels[x][y])

			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				image[newX][newY] = newColor
			}
		}
	}
	global.ExecutionTime.SetText(time.Since(start).String())
	return image
}

func singleMultithread(fun PixelTransformFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixel[x][y])
			}
		}
	})
	return image
}

func bothMultithread(fun PixelsTransformFunc, width, height int, pixelOne, pixelTwo, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixelOne[x][y], pixelTwo[x][y])
			}
		}
	})
	return image
}

func axisMultithread(fun AxisTransformFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
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

func Process(service interface{}) func() {
	return func() {
		if *global.ImageOne == nil && *global.ImageTwo == nil {
			slog.Error("No image set")
			return
		}

		var pixelsTwo *[][]color.RGBA
		imgOne := *global.ImageOne
		if imgOne == nil {
			imgOne = *global.ImageTwo
		}

		pixelsOne := ConvertImageToPixels(imgOne)
		if *global.ImageTwo != nil {
			p := ConvertImageToPixels(*global.ImageTwo)
			pixelsTwo = &p
		}

		xLen := len(pixelsOne)
		yLen := len(pixelsOne[0])

		newImage := make([][]color.RGBA, xLen)
		for i := range xLen {
			newImage[i] = make([]color.RGBA, yLen)
		}

		switch s := service.(type) {
		case PixelTransformFunc:
			if global.UseSingleThread {
				newImage = single(s, xLen, yLen, pixelsOne, newImage)
				break
			}
			newImage = singleMultithread(s, xLen, yLen, pixelsOne, newImage)
		case PixelsTransformFunc:
			if pixelsTwo != nil {
				if global.UseSingleThread {
					newImage = both(s, xLen, yLen, pixelsOne, *pixelsTwo, newImage)
				}
				newImage = bothMultithread(s, xLen, yLen, pixelsOne, *pixelsTwo, newImage)
			}
		case AxisTransformFunc:
			if global.UseSingleThread {
				newImage = axis(s, xLen, yLen, pixelsOne, newImage)
			}
			newImage = axisMultithread(s, xLen, yLen, pixelsOne, newImage)
		default:
			slog.Error("invalid service", "type", s)
			return
		}

		global.FinalImage.Image = ConvertPixelsToImage(newImage)
		global.FinalImage.Refresh()
	}
}
