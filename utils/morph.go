package utils

import (
	"image/color"
)

type MorphOpsFunc func([][]color.RGBA) [][]color.RGBA

func morph(fun MorphOpsFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	return fun(pixel)
}
