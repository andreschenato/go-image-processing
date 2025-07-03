package service

import (
	"image/color"
	"image_processing/utils"
)

func Dilation() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		mask := [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}

		maxR := 0
		maxG := 0
		maxB := 0

		for x, row := range pixels {
			for y, px := range row {
				if mask[x][y] == 1 {
					maxR = max(maxR, int(px.R))
					maxG = max(maxG, int(px.G))
					maxB = max(maxB, int(px.B))
				}
			}
		}

		return color.RGBA{
			R: uint8(maxR),
			G: uint8(maxG),
			B: uint8(maxB),
			A: 255,
		}
	}
}

func Erosion() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		mask := [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}

		minR := 255
		minG := 255
		minB := 255

		for x, row := range pixels {
			for y, px := range row {
				if mask[x][y] == 1 {
					minR = min(minR, int(px.R))
					minG = min(minG, int(px.G))
					minB = min(minB, int(px.B))
				}
			}
		}

		return color.RGBA{
			R: uint8(minR),
			G: uint8(minG),
			B: uint8(minB),
			A: 255,
		}
	}
}
