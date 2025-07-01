package utils

import (
	"image"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type HistEqualizationFunc func(color.RGBA) color.RGBA

func HistogramValues(img image.Image) (image.Image, []int) {
	pixels := ConvertImageToPixels(img)
	width := len(pixels)
	height := len(pixels[0])

	hist := make([]int, 256)
	for x := range width {
		for y := range height {
			p := pixels[x][y]
			gray := uint8((float64(p.R) + float64(p.G) + float64(p.B)) / 3)
			hist[gray]++
		}
	}

	pts := make(plotter.XYs, 256)
	for i := range 256 {
		pts[i].X = float64(i)
		pts[i].Y = float64(hist[i])
	}

	p := plot.New()
	p.HideAxes()
	p.BackgroundColor = color.RGBA{40, 40, 42, 255}
	p.X.Min = 0
	p.X.Max = 255

	bars, err := plotter.NewHistogram(pts, 256)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Color = color.White
	bars.Width = 1
	bars.Color = color.White

	p.Add(bars)

	img_canvas := vgimg.New(800, 250)

	dc := draw.New(img_canvas)
	p.Draw(dc)

	resultImg := img_canvas.Image()

	return resultImg, hist
}

func histogramEqualization(fun HistEqualizationFunc, width, height int, pixel, image [][]color.RGBA) [][]color.RGBA {
	_, histogram := HistogramValues(ConvertPixelsToImage(pixel))

	cdf := make([]int, 256)
	cdf[0] = histogram[0]
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + histogram[i]
	}

	cdfMin := 0
	for i := range 256 {
		if cdf[i] > 0 {
			cdfMin = cdf[i]
			break
		}
	}

	lookupTable := make([]uint8, 256)
	for i := range 256 {
		numerator := float64(cdf[i] - cdfMin)
		denominator := float64((width * height) - cdfMin)
		result := numerator / denominator * 255.0
		lookupTable[i] = uint8(math.Max(0, math.Min(255, math.Floor(result))))
	}

	parallelizeProcessing(width, height, func(startX, endX, height int) {
		for x := startX; x < endX; x++ {
			for y := range height {
				p := pixel[x][y]
				gray := uint8((float64(p.R) + float64(p.G) + float64(p.B)) / 3)
				newGray := lookupTable[gray]

				equalizedColor := fun(color.RGBA{R: newGray, G: newGray, B: newGray, A: 255})
				image[x][y] = equalizedColor
			}
		}
	})

	return image
}
