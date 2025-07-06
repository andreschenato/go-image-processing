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

	global.FinalHist = canvas.NewImageFromImage(nil)
	global.FinalHist.FillMode = canvas.ImageFillStretch
	global.FinalHist.SetMinSize(fyne.NewSize(250, 250))

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
			global.FinalHist.Image = nil
			global.FinalImage.Refresh()
			global.FinalHist.Refresh()
		}),
	)

	imgContainer := container.NewStack(placeholder, global.FinalImage)
	histContainer := container.NewStack(placeholder, global.FinalHist)

	return container.NewVBox(
		container.NewPadded(imgContainer),
		container.NewPadded(histContainer),
		container.NewPadded(downloadBtn),
		container.NewPadded(clearBtn),
	)
}
