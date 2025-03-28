package utils

import (
	"image"
	"image/color"
)

func ConvertImageToPixels(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	var pixels [][]color.Color

	for i := range size.X {
		var y []color.Color
		for j := range size.Y {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}

	return pixels
}

func ConvertPixelsToImage(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := range len(pixels) {
		for y := range len(pixels[0]) {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}

	return nImg
}
