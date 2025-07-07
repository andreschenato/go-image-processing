package utils

import (
	"image/color"
	"image_processing/global"
	"log/slog"
)

func Process(service interface{}) func() {
	return func() {
		if *global.ImageOne == nil && *global.ImageTwo == nil {
			label := "No image set"
			slog.Error(label)
			warning(label)
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
			if *global.ImageOne == nil || *global.ImageTwo == nil {
				label := "This kind of function needs two images"
				slog.Error(label)
				warning(label)
				return
			}
			if len(pixelsOne) != len(*pixelsTwo) {
				label := "The two images need to have the same size"
				slog.Error(label)
				warning(label)
				return
			}
			newImage = both(s, width, height, pixelsOne, *pixelsTwo, newImage)
		case AxisTransformFunc:
			newImage = axis(s, width, height, pixelsOne, newImage)
		case HistEqualizationFunc:
			newImage = histogramEqualization(s, width, height, pixelsOne, newImage)
		case LowPassFilterFunc:
			newImage = lowPass(s, width, height, pixelsOne, newImage)
		case HighPassFilterFunc:
			newImage = highPass(s, width, height, pixelsOne, newImage)
		case MorphOpsFunc:
			newImage = Morph(s, width, height, pixelsOne, newImage)
		case ComboMorphOpsFunc:
			newImage = comboMorph(s, width, height, pixelsOne)
		default:
			slog.Error("invalid service", "type", s)
			return
		}

		global.FinalImage.Image = ConvertPixelsToImage(newImage)
		global.FinalImage.Refresh()
		global.FinalHist.Image, _ = HistogramValues(global.FinalImage.Image)
		global.FinalHist.Refresh()
	}
}
