package global

import (
	"image"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var (
	ImageOne *image.Image
	ImageTwo *image.Image

	FinalImage *canvas.Image

	ExecutionTime widget.Label = *widget.NewLabel("0s")
	UseSingleThread bool = false
)