package components

import (
	"image"
	"image/color"
	"image_processing/global"
	"image_processing/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ImageUploadView() (*fyne.Container, *image.Image) {
	placeholder := canvas.NewRectangle(color.RGBA{40, 40, 42, 255})
	placeholder.SetMinSize(fyne.NewSize(250, 250))

	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(250, 250))

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(250, 50))

	uploadBtn := container.New(
		fixedLayout,
		widget.NewButton("Upload Image", func() {
			utils.UploadImage(img)
			global.Hist.Image, _ = utils.HistogramValues(img.Image)
			global.Hist.Refresh()
		}),
	)

	clearBtn := container.New(
		fixedLayout,
		widget.NewButton("Clear", func() {
			img.Image = nil
			img.Refresh()
		}),
	)

	imgContainer := container.NewStack(placeholder, img)

	imgCombo := container.NewVBox(
		container.NewPadded(imgContainer),
		container.NewPadded(uploadBtn),
		container.NewPadded(clearBtn),
	)
	return container.NewCenter(imgCombo), &img.Image
}
