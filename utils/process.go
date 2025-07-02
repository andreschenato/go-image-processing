package utils

import (
	"image/color"
	"image_processing/global"
	"log/slog"
)

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

		width := len(pixelsOne)
		height := len(pixelsOne[0])

		newImage := make([][]color.RGBA, width)
		for i := range width {
			newImage[i] = make([]color.RGBA, height)
		}

		switch s := service.(type) {
		case PixelTransformFunc:
			newImage = single(s, width, height, pixelsOne, newImage)
		case PixelsTransformFunc:
			if pixelsTwo == nil {
				return
			}
			newImage = both(s, width, height, pixelsOne, *pixelsTwo, newImage)
		case AxisTransformFunc:
			newImage = axis(s, width, height, pixelsOne, newImage)
		case HistEqualizationFunc:
			newImage = histogramEqualization(s, width, height, pixelsOne, newImage)
		case LowPassFilterFunc:
			newImage = lowPass(s, width, height, pixelsOne, newImage)
		default:
			slog.Error("invalid service", "type", s)
			return
		}

		global.FinalImage.Image = ConvertPixelsToImage(newImage)
		global.FinalImage.Refresh()
		global.Hist.Image, _ = HistogramValues(global.FinalImage.Image)
		global.Hist.Refresh()
	}
}
