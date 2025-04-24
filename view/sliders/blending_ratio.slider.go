package sliders

import (
	"image/color"
	"image_processing/global"
	"image_processing/service"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BlendingRatioSlider() *fyne.Container {
	slider := widget.NewSlider(0, 1)
	slider.SetValue(1)
	slider.Step = 0.01

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(500, 10))

	txt := widget.NewLabel(strconv.FormatFloat(slider.Value, 'f', -1, 32))

	sliderContainer := container.NewVBox(
		widget.NewLabel("Blending"),
		container.New(fixedLayout, slider),
		txt,
	)

	box := canvas.NewRectangle(color.RGBA{40,40,42,255})
	box.SetMinSize(
		fyne.NewSize(
			slider.MinSize().Width + 10,
			slider.MinSize().Height + 10,
		),
	)

	sliderBox := container.NewStack(
		box,
		sliderContainer,
	)

	slider.OnChanged = func(value float64) {
		txt.SetText(strconv.FormatFloat(slider.Value, 'f', -1, 32))
		if *global.ImageOne != nil && *global.ImageTwo != nil {
			global.FinalImage.Image = *service.BlendImages(value)
			global.FinalImage.Refresh()
		}
	}

	return container.NewCenter(container.NewHBox(sliderBox))
}
