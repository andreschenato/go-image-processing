package service

import (
	"image/color"
	"image_processing/utils"
)

func Horizontal() utils.AxisTransformFunc {
	return func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA) {
		newX := width - 1 - x
		return newX, y, pixel
	}
}

func Vertical() utils.AxisTransformFunc {
	return func(x, y, width, height int, pixel color.RGBA) (int, int, color.RGBA) {
		newY := height - 1 - y
		return x, newY, pixel
	}
}
