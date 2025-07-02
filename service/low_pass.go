package service

import (
	"image/color"
	"image_processing/utils"
)

func Min() utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		minR := uint8(255)
		minG := uint8(255)
		minB := uint8(255)

		for _, row := range pixels {
			for _, px := range row {
				minR = min(px.R, minR)
				minG = min(px.G, minG)
				minB = min(px.B, minB)
			}
		}

		return color.RGBA{
			R: minR,
			G: minG,
			B: minB,
			A: 255,
		}
	}
}

func Max() utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		maxR := uint8(0)
		maxG := uint8(0)
		maxB := uint8(0)

		for _, row := range pixels {
			for _, px := range row {
				maxR = max(px.R, maxR)
				maxG = max(px.G, maxG)
				maxB = max(px.B, maxB)
			}
		}

		return color.RGBA{
			R: maxR,
			G: maxG,
			B: maxB,
			A: 255,
		}
	}
}
func Mean() utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var sumR, sumG, sumB int

		for _, row := range pixels {
			for _, px := range row {
				sumR += int(px.R)
				sumG += int(px.G)
				sumB += int(px.B)
			}
		}

		return color.RGBA{
			R: uint8(sumR / 9),
			G: uint8(sumG / 9),
			B: uint8(sumB / 9),
			A: 255,
		}
	}
}
