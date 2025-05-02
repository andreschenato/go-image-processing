package utils

import (
	"image"
	"image/color"
)

func ConvertImageToPixels(img image.Image) [][]color.RGBA {
	size := img.Bounds().Size()
	var pixels [][]color.RGBA

	for i := range size.X {
		var y []color.RGBA
		for j := range size.Y {
			y = append(y, color.RGBAModel.Convert(img.At(i, j)).(color.RGBA))
		}
		pixels = append(pixels, y)
	}

	return pixels
}

func ConvertPixelsToImage(pixels [][]color.RGBA) image.Image {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := range len(pixels) {
		for y := range len(pixels[0]) {
			nImg.Set(x, y, pixels[x][y])
		}
	}

	return nImg
}
