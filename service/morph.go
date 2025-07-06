package service

import (
	"image/color"
	"image_processing/global"
	"image_processing/utils"
)

func buildSquareMask() [][]int {
	mask := make([][]int, global.MaskSize)
	for i := range global.MaskSize {
		mask[i] = make([]int, global.MaskSize)
		for j := range global.MaskSize {
			mask[i][j] = 1
		}
	}
	return mask
}

func applyMorph(pixels [][]color.RGBA, mask [][]int, isDilation bool) [][]color.RGBA {
	height := len(pixels)
	width := len(pixels[0])
	maskSize := global.MaskSize
	offset := maskSize / 2

	result := make([][]color.RGBA, height)
	for i := range result {
		result[i] = make([]color.RGBA, width)
	}

	for i := range height {
		for j := range width {
			var R, G, B int
			if isDilation {
				R, G, B = 0, 0, 0
			} else {
				R, G, B = 255, 255, 255
			}

			for mx := range maskSize {
				for my := range maskSize {
					ix := i + mx - offset
					iy := j + my - offset
					if ix >= 0 && ix < height && iy >= 0 && iy < width && mask[mx][my] == 1 {
						px := pixels[ix][iy]
						if isDilation {
							R = max(R, int(px.R))
							G = max(G, int(px.G))
							B = max(B, int(px.B))
						} else {
							R = min(R, int(px.R))
							G = min(G, int(px.G))
							B = min(B, int(px.B))
						}
					}
				}
			}

			result[i][j] = color.RGBA{
				R: uint8(R),
				G: uint8(G),
				B: uint8(B),
				A: 255,
			}
		}
	}

	return result
}

func Dilation() utils.MorphOpsFunc {
	return func(pixels [][]color.RGBA) [][]color.RGBA {
		mask := buildSquareMask()
		return applyMorph(pixels, mask, true)
	}
}

func Erosion() utils.MorphOpsFunc {
	return func(pixels [][]color.RGBA) [][]color.RGBA {
		mask := buildSquareMask()
		return applyMorph(pixels, mask, false)
	}
}

func Opening() utils.MorphOpsFunc {
	return func(pixels [][]color.RGBA) [][]color.RGBA {
		mask := buildSquareMask()
		eroded := applyMorph(pixels, mask, false)
		return applyMorph(eroded, mask, true)
	}
}

func Closing() utils.MorphOpsFunc {
	return func(pixels [][]color.RGBA) [][]color.RGBA {
		mask := buildSquareMask()
		dilated := applyMorph(pixels, mask, true)
		return applyMorph(dilated, mask, false)
	}
}

func Outline() utils.MorphOpsFunc {
	return func(pixels [][]color.RGBA) [][]color.RGBA {
		mask := buildSquareMask()
		eroded := applyMorph(pixels, mask, false)
		for x, row := range pixels {
			for y, px := range row {
				pixels[x][y] = Subtract()(px, eroded[x][y])
			}
		}
		return pixels
	}
}
