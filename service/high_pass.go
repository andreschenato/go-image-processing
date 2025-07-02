package service

import (
	"image/color"
	"image_processing/utils"
	"math"
)

func Sobel() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var rX, gX, bX int
		var rY, gY, bY int

		maskX := [][]int{
			{1, 0, -1},
			{2, 0, -2},
			{1, 0, -1},
		}

		maskY := [][]int{
			{1, 2, 1},
			{0, 0, 0},
			{-1, -2, -1},
		}

		for x, row := range pixels {
			for y, px := range row {
				rX += maskX[x][y] * int(px.R)
				gX += maskX[x][y] * int(px.G)
				bX += maskX[x][y] * int(px.B)
			}
		}

		for x, row := range pixels {
			for y, px := range row {
				rY += maskY[x][y] * int(px.R)
				gY += maskY[x][y] * int(px.G)
				bY += maskY[x][y] * int(px.B)
			}
		}

		rX = min(rX, 255)
		gX = min(gX, 255)
		bX = min(bX, 255)
		rY = min(rY, 255)
		gY = min(gY, 255)
		bY = min(bY, 255)

		r := uint8(math.Sqrt(float64(int(math.Pow(float64(rX), 2)) + int(math.Pow(float64(rY), 2)))))
		g := uint8(math.Sqrt(float64(int(math.Pow(float64(gX), 2)) + int(math.Pow(float64(gY), 2)))))
		b := uint8(math.Sqrt(float64(int(math.Pow(float64(bX), 2)) + int(math.Pow(float64(bY), 2)))))

		return color.RGBA{
			R: uint8(min(r, 255)),
			G: uint8(min(g, 255)),
			B: uint8(min(b, 255)),
			A: 255,
		}
	}
}

func Prewitt() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var rX, gX, bX int
		var rY, gY, bY int

		maskX := [][]int{
			{1, 0, -1},
			{1, 0, -1},
			{1, 0, -1},
		}

		maskY := [][]int{
			{1, 1, 1},
			{0, 0, 0},
			{-1, -1, -1},
		}

		for x, row := range pixels {
			for y, px := range row {
				rX += maskX[x][y] * int(px.R)
				gX += maskX[x][y] * int(px.G)
				bX += maskX[x][y] * int(px.B)
			}
		}

		for x, row := range pixels {
			for y, px := range row {
				rY += maskY[x][y] * int(px.R)
				gY += maskY[x][y] * int(px.G)
				bY += maskY[x][y] * int(px.B)
			}
		}

		rX = min(rX, 255)
		gX = min(gX, 255)
		bX = min(bX, 255)
		rY = min(rY, 255)
		gY = min(gY, 255)
		bY = min(bY, 255)

		r := uint8(math.Sqrt(float64(int(math.Pow(float64(rX), 2)) + int(math.Pow(float64(rY), 2)))))
		g := uint8(math.Sqrt(float64(int(math.Pow(float64(gX), 2)) + int(math.Pow(float64(gY), 2)))))
		b := uint8(math.Sqrt(float64(int(math.Pow(float64(bX), 2)) + int(math.Pow(float64(bY), 2)))))

		return color.RGBA{
			R: uint8(min(r, 255)),
			G: uint8(min(g, 255)),
			B: uint8(min(b, 255)),
			A: 255,
		}
	}
}

func LaplacianExterior() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var rX, gX, bX int

		maskX := [][]int{
			{1, 1, 1},
			{1, -8, 1},
			{1, 1, 1},
		}

		for x, row := range pixels {
			for y, px := range row {
				rX += maskX[x][y] * int(px.R)
				gX += maskX[x][y] * int(px.G)
				bX += maskX[x][y] * int(px.B)
			}
		}

		rX = max(min(rX, 255), 0)
		gX = max(min(gX, 255), 0)
		bX = max(min(bX, 255), 0)

		return color.RGBA{
			R: uint8(min(rX, 255)),
			G: uint8(min(gX, 255)),
			B: uint8(min(bX, 255)),
			A: 255,
		}
	}
}

func LaplacianInterior() utils.HighPassFilterFunc {
	return func(pixels [][]color.RGBA) color.RGBA {
		var rX, gX, bX int

		maskX := [][]int{
			{-1, -1, -1},
			{-1, 8, -1},
			{-1, -1, -1},
		}

		for x, row := range pixels {
			for y, px := range row {
				rX += maskX[x][y] * int(px.R)
				gX += maskX[x][y] * int(px.G)
				bX += maskX[x][y] * int(px.B)
			}
		}

		rX = max(min(rX, 255), 0)
		gX = max(min(gX, 255), 0)
		bX = max(min(bX, 255), 0)

		return color.RGBA{
			R: uint8(min(rX, 255)),
			G: uint8(min(gX, 255)),
			B: uint8(min(bX, 255)),
			A: 255,
		}
	}
}
