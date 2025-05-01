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

func single(fun PixelTransformFunc) {
	start := time.Now()
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

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

					newColor := fun(originalColor)

					newImage[x][y] = newColor
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()

	global.ExecutionTime.SetText(time.Since(start).String())
}

func both(fun PixelsTransformFunc) {
	pixelsOne := ConvertImageToPixels(*global.ImageOne)
	pixelsTwo := ConvertImageToPixels(*global.ImageTwo)

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
					ogColorOne, ok := color.RGBAModel.Convert(pixelsOne[x][y]).(color.RGBA)
					if !ok {
						slog.Error("conversion went wrong")
					}

					ogColorTwo, ok := color.RGBAModel.Convert(pixelsTwo[x][y]).(color.RGBA)
					if !ok {
						slog.Error("conversion went wrong")
					}

					newColor := fun(ogColorOne, ogColorTwo)

					newImage[x][y] = newColor
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
}

func axis(fun AxisTransformFunc) {
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

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

					newX, newY, newColor := fun(x, y, xLen, yLen, originalColor)

					if newX >= 0 && newX < xLen && newY >= 0 && newY < yLen {
						newImage[newX][newY] = newColor
					}
				}
			}
		}(startX, endX)
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
}

func threadOne(fun PixelTransformFunc) {
	start := time.Now()
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

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
					slog.Error("type conversion went wrong")
				}
				newColor := fun(originalColor)

				newImage[x][y] = newColor
				wg.Done()
			}(x, y)
		}
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
	global.ExecutionTime.SetText(time.Since(start).String())
}

func threadTwo(fun PixelsTransformFunc) {
	start := time.Now()

	pixelsOne := ConvertImageToPixels(*global.ImageOne)
	pixelsTwo := ConvertImageToPixels(*global.ImageTwo)

	xLen := len(pixelsOne)
	yLen := len(pixelsOne[0])

	newImage := make([][]color.Color, xLen)
	for i := range newImage {
		newImage[i] = make([]color.Color, yLen)
	}

	wg := sync.WaitGroup{}
	for x := range xLen {
		for y := range yLen {
			wg.Add(1)
			go func(x, y int) {
				ogColorOne, ok := color.RGBAModel.Convert(pixelsOne[x][y]).(color.RGBA)
				if !ok {
					slog.Error("conversion went wrong")
				}

				ogColorTwo, ok := color.RGBAModel.Convert(pixelsTwo[x][y]).(color.RGBA)
				if !ok {
					slog.Error("conversion went wrong")
				}

				newColor := fun(ogColorOne, ogColorTwo)

				newImage[x][y] = newColor
				wg.Done()
			}(x, y)
		}
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
	global.ExecutionTime.SetText(time.Since(start).String())
}

func threadAxis(fun AxisTransformFunc) {
	start := time.Now()
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}

	pixels := ConvertImageToPixels(img)

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
					slog.Error("type conversion went wrong")
				}

				newX, newY, newColor := fun(x, y, xLen, yLen, originalColor)

				if newX >= 0 && newX < xLen && newY >= 0 && newY < yLen {
					newImage[newX][newY] = newColor
				}
				wg.Done()
			}(x, y)
		}
	}
	wg.Wait()

	global.FinalImage.Image = ConvertPixelsToImage(newImage)
	global.FinalImage.Refresh()
	global.ExecutionTime.SetText(time.Since(start).String())
}

func Process(service interface{}) func() {
	return func() {
		switch s := service.(type) {
		case PixelTransformFunc:
			if *global.ImageOne != nil || *global.ImageTwo != nil {
				if global.UseSingleThread {
					threadOne(s)
					return
				}
				single(s)
				return
			}
			slog.Error("No image set")
			return
		case PixelsTransformFunc:
			if *global.ImageOne != nil && *global.ImageTwo != nil {
				if global.UseSingleThread {
					threadTwo(s)
					return
				}
				both(s)
				return
			}
			slog.Error("Images aren't set")
			return
		case AxisTransformFunc:
			if *global.ImageOne != nil || *global.ImageTwo != nil {
				if global.UseSingleThread {
					threadAxis(s)
					return
				}
				axis(s)
				return
			}
			slog.Error("No image set")
			return
		default:
			slog.Error("invalid service", "type", s)
			return
		}
	}
}
