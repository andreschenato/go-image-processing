package service

import (
	"image/color"
	"image_processing/utils"
)

func Dilation() utils.MorphOpsFunc {
	return func(pixels []color.RGBA, mask [][]int) color.RGBA {
		var R, G, B int = 0, 0, 0

		for _, px := range pixels {
			R = max(R, int(px.R))
			G = max(G, int(px.G))
			B = max(B, int(px.B))
		}

		return color.RGBA{
			R: uint8(R),
			G: uint8(G),
			B: uint8(B),
			A: 255,
		}
	}
}

func Erosion() utils.MorphOpsFunc {
	return func(pixels []color.RGBA, mask [][]int) color.RGBA {
		var R, G, B int = 255, 255, 255

		for _, px := range pixels {
			R = min(R, int(px.R))
			G = min(G, int(px.G))
			B = min(B, int(px.B))
		}

		return color.RGBA{
			R: uint8(R),
			G: uint8(G),
			B: uint8(B),
			A: 255,
		}
	}
}

func Opening() utils.ComboMorphOpsFunc {
	return func(width, height int, pixels [][]color.RGBA) [][]color.RGBA {
		eroded := make([][]color.RGBA, height)
		for i := range eroded {
			eroded[i] = make([]color.RGBA, width)
		}

		result := make([][]color.RGBA, height)
		for i := range result {
			result[i] = make([]color.RGBA, width)
		}

		eroded = utils.Morph(Erosion(), width, height, pixels, eroded)
		result = utils.Morph(Dilation(), width, height, eroded, result)

		return result
	}
}

func Closing() utils.ComboMorphOpsFunc {
	return func(width, height int, pixels [][]color.RGBA) [][]color.RGBA {
		dilated := make([][]color.RGBA, height)
		for i := range dilated {
			dilated[i] = make([]color.RGBA, width)
		}

		result := make([][]color.RGBA, height)
		for i := range result {
			result[i] = make([]color.RGBA, width)
		}

		dilated = utils.Morph(Dilation(), width, height, pixels, dilated)
		result = utils.Morph(Erosion(), width, height, dilated, result)

		return result
	}
}

func Outline() utils.ComboMorphOpsFunc {
	return func(width, height int, pixels [][]color.RGBA) [][]color.RGBA {
		eroded := make([][]color.RGBA, height)
		for i := range eroded {
			eroded[i] = make([]color.RGBA, width)
		}

		eroded = utils.Morph(Erosion(), width, height, pixels, eroded)
		for x, row := range pixels {
			for y, px := range row {
				pixels[x][y] = Subtract()(px, eroded[x][y])
			}
		}
		return pixels
	}
}
