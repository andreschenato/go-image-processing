package components

import (
	"image/color"
	"image_processing/global"
	"image_processing/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ProcessedImage() *fyne.Container {
	placeholder := canvas.NewRectangle(color.RGBA{40, 40, 42, 255})
	placeholder.SetMinSize(fyne.NewSize(250, 250))

	global.FinalImage = canvas.NewImageFromImage(nil)
	global.FinalImage.FillMode = canvas.ImageFillContain
	global.FinalImage.SetMinSize(fyne.NewSize(250, 250))

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(250, 50))

	downloadBtn := container.New(
		fixedLayout,
		widget.NewButton("Download Image", func() {
			utils.DownloadImage()
		}),
	)

	clearBtn := container.New(
		fixedLayout,
		widget.NewButton("Clear", func() {
			global.FinalImage.Image = nil
			global.FinalImage.Refresh()
		}),
	)

	imgContainer := container.NewStack(placeholder, global.FinalImage)

	return container.NewVBox(
		container.NewPadded(imgContainer),
		container.NewPadded(downloadBtn),
		container.NewPadded(clearBtn),
	)
}
