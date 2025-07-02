package service

import (
	"image/color"
	"image_processing/global"
	"image_processing/utils"
	"math"
	"slices"
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

func Conservative() utils.LowPassFilterFunc {
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

		median := count / 2

		mainR := valR[median]
		mainG := valG[median]
		mainB := valB[median]

		valR = slices.Delete(valR, median, median+1)
		valG = slices.Delete(valG, median, median+1)
		valB = slices.Delete(valB, median, median+1)

		if mainR > slices.Max(valR) {
			mainR = slices.Max(valR)
		} else if mainR < slices.Min(valR) {
			mainR = slices.Min(valR)
		}

		if mainG > slices.Max(valG) {
			mainG = slices.Max(valG)
		} else if mainG < slices.Min(valG) {
			mainG = slices.Min(valG)
		}

		if mainB > slices.Max(valB) {
			mainB = slices.Max(valB)
		} else if mainB < slices.Min(valB) {
			mainB = slices.Min(valB)
		}

		return color.RGBA{
			R: uint8(mainR),
			G: uint8(mainG),
			B: uint8(mainB),
			A: 255,
		}
	}
}

func Gaussian(sigma float64) utils.LowPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		center := int(global.MaskSize / 2)

		weights := make([][]float64, global.MaskSize)
		var totalKernelWeight float64

		for i := range global.MaskSize {
			weights[i] = make([]float64, global.MaskSize)
			for j := range global.MaskSize {
				x := float64(j - center)
				y := float64(i - center)

				exponent := -(math.Pow(x, 2) + math.Pow(y, 2)) / (2 * math.Pow(sigma, 2))
				weight := (1.0 / (2 * math.Pi * math.Pow(sigma, 2))) * math.Exp(exponent)

				weights[i][j] = weight
				totalKernelWeight += weight
			}
		}

		for i := range global.MaskSize {
			for j := range global.MaskSize {
				weights[i][j] /= totalKernelWeight
			}
		}

		var weightedR, weightedG, weightedB float64

		for i, row := range pixels {
			for j, px := range row {
				normalizedWeight := weights[i][j]
				weightedR += float64(px.R) * normalizedWeight
				weightedG += float64(px.G) * normalizedWeight
				weightedB += float64(px.B) * normalizedWeight
			}
		}

		r := math.Max(0, math.Min(255, weightedR))
		g := math.Max(0, math.Min(255, weightedG))
		b := math.Max(0, math.Min(255, weightedB))

		return color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 255,
		}
	}
}
