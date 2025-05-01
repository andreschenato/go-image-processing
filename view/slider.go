package view

import (
	"image/color"
	"image_processing/utils"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Slider[T any](min, max, initial, step float64, label string, service func(value float64) T) *fyne.Container {
	slider := widget.NewSlider(min, max)
	slider.SetValue(initial)
	slider.Step = step

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(500, 10))

	txt := widget.NewLabel(strconv.FormatFloat(slider.Value, 'f', -1, 32))

	sliderContainer := container.NewVBox(
		widget.NewLabel(label),
		container.New(fixedLayout, slider),
		txt,
	)

	box := canvas.NewRectangle(color.RGBA{40, 40, 42, 255})
	box.SetMinSize(
		fyne.NewSize(
			slider.MinSize().Width+10,
			slider.MinSize().Height+10,
		),
	)

	sliderBox := container.NewStack(
		box,
		sliderContainer,
	)

	slider.OnChanged = func(value float64) {
		txt.SetText(strconv.FormatFloat(slider.Value, 'f', -1, 32))
		utils.Process(service(value))()
	}

	return container.NewPadded(
		container.NewCenter(
			container.NewHBox(sliderBox),
		),
	)
}
