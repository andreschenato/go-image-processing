package global

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var (
	ImageOne *image.Image
	ImageTwo *image.Image

	FinalImage *canvas.Image

	FinalHist *canvas.Image

	MaskSize int = 3

	UseDiamondMask bool = false

	Window fyne.Window
)