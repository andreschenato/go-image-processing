package service

import (
	"image/color"
	"image_processing/utils"
	"sort"
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
		count := 0

		for _, row := range pixels {
			for _, px := range row {
				sumR += int(px.R)
				sumG += int(px.G)
				sumB += int(px.B)
				count++
			}
		}

		return color.RGBA{
			R: uint8(sumR / count),
			G: uint8(sumG / count),
			B: uint8(sumB / count),
			A: 255,
		}
	}
}

func Median() utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var valR, valG, valB []int
		count := 0

		for _, row := range pixels {
			for _, px := range row {
				valR = append(valR, int(px.R))
				valG = append(valG, int(px.G))
				valB = append(valB, int(px.B))
				count++
			}
		}

		sort.Ints(valR)
		sort.Ints(valG)
		sort.Ints(valB)

		median := count / 2

		return color.RGBA{
			R: uint8(valR[median]),
			G: uint8(valG[median]),
			B: uint8(valB[median]),
			A: 255,
		}
	}
}

func Order(target float64) utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var valR, valG, valB []int
		count := 0

		for _, row := range pixels {
			for _, px := range row {
				valR = append(valR, int(px.R))
				valG = append(valG, int(px.G))
				valB = append(valB, int(px.B))
				count++
			}
		}

		sort.Ints(valR)
		sort.Ints(valG)
		sort.Ints(valB)

		t := int(target)

		return color.RGBA{
			R: uint8(valR[t]),
			G: uint8(valG[t]),
			B: uint8(valB[t]),
			A: 255,
		}
	}
}
