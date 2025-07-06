package utils

import (
	"image/color"
	"image_processing/global"
)

type MorphOpsFunc func([]color.RGBA, [][]int) color.RGBA
type ComboMorphOpsFunc func(int, int, [][]color.RGBA) [][]color.RGBA

func Morph(fun MorphOpsFunc, width, height int, pixels, image [][]color.RGBA) [][]color.RGBA {
	mask := buildSquareMask()
	if global.UseDiamondMask {
		mask = buildDiamondMask()
	}

	maskSize := global.MaskSize
	offset := maskSize / 2

	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for i := range height {
			for j := startX; j < endX; j++ {
				var validPixels []color.RGBA

				for mx := range maskSize {
					for my := range maskSize {
						ix := i + mx - offset
						iy := j + my - offset

						if ix >= 0 && ix < height && iy >= 0 && iy < width && mask[mx][my] == 1 {
							validPixels = append(validPixels, pixels[ix][iy])
						}
					}
				}

				if len(validPixels) > 0 {
					image[i][j] = fun(validPixels, mask)
				} else {
					image[i][j] = pixels[i][j]
				}
			}
		}
	})

	return image
}

func comboMorph(fun ComboMorphOpsFunc, width, height int, pixels [][]color.RGBA) [][]color.RGBA {
	return fun(width, height, pixels)
}
