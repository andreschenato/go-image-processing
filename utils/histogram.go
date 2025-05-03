package utils

import (
	"image"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func HistogramValues(img image.Image) image.Image {
	var hist [256]int

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((float64(r) + float64(g) + float64(b)) / 3)
			hist[gray]++
		}
	}

	pts := make(plotter.XYs, 256)
	for i := range 256 {
		pts[i].X = float64(i)
		pts[i].Y = float64(hist[i])
	}

	p := plot.New()
	p.Title.Text = "Histograma"
	p.X.Min = 0
	p.X.Max = 255

	bars, err := plotter.NewHistogram(pts, 256)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Color = color.Black
	bars.Width = 1
	bars.Color = color.Black

	p.Add(bars)

	img_canvas := vgimg.New(800, 250)

	dc := draw.New(img_canvas)
	p.Draw(dc)

	resultImg := img_canvas.Image()

	return resultImg
}
