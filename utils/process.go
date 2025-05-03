package utils

import (
	"image"
	"image/color"
	"image_processing/global"
	"log/slog"
	"math"
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
	slog.Info(time.Since(start).String())
}

func single(fun PixelTransformFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixel[x][y])
			}
		}
	})
	return image
}

func both(fun PixelsTransformFunc, width, height int, pixelOne, pixelTwo, image [][]color.RGBA) [][]color.RGBA {
	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				image[x][y] = fun(pixelOne[x][y], pixelTwo[x][y])
			}
		}
	})
	return image
}

func axis(fun AxisTransformFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
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
			newImage = single(s, xLen, yLen, pixelsOne, newImage)
		case PixelsTransformFunc:
			if pixelsTwo == nil {
				return
			}
			newImage = both(s, xLen, yLen, pixelsOne, *pixelsTwo, newImage)
		case AxisTransformFunc:
			newImage = axis(s, xLen, yLen, pixelsOne, newImage)
		default:
			slog.Error("invalid service", "type", s)
			return
		}

		global.FinalImage.Image = ConvertPixelsToImage(newImage)
		global.FinalImage.Refresh()
		global.Hist.Image = HistogramValues(global.FinalImage.Image)
		global.Hist.Refresh()
	}
}

func HistogramEqualization() {
	img := *global.ImageOne
	if img == nil {
		img = *global.ImageTwo
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	totalPixels := width * height

	var histogram [256]int

	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((float64(r) + float64(g) + float64(b)) / 3)
			histogram[gray]++
		}
	}

	var cdf [256]int
	cdf[0] = histogram[0]
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + histogram[i]
	}

	cdfMin := 0
	for i := range 256 {
		if cdf[i] > 0 {
			cdfMin = cdf[i]
			break
		}
	}

	var lookupTable [256]uint8
	for i := range 256 {
		if cdf[i] == 0 {
			lookupTable[i] = 0
		} else {
			numerator := float64(cdf[i] - cdfMin)
			denominator := float64(totalPixels - cdfMin)
			result := numerator / denominator * 255.0
			lookupTable[i] = uint8(math.Floor(result))
		}
	}

	equalizedImg := image.NewRGBA(bounds)
	for y := range height {
		for x := range width {
			r, g, b, a := img.At(x, y).RGBA()
			gray := uint8((float64(r) + float64(g) + float64(b)) / 3)

			newGray := lookupTable[gray]

			equalizedImg.Set(x, y, color.RGBA{
				R: newGray,
				G: newGray,
				B: newGray,
				A: uint8(a >> 8),
			})
		}
	}

	global.FinalImage.Image = equalizedImg
	global.FinalImage.Refresh()
	global.Hist.Image = HistogramValues(global.FinalImage.Image)
	global.Hist.Refresh()
}
