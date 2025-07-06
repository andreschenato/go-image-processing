package generics

import (
	"image/color"
	"image_processing/global"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MaskSlider() *fyne.Container {
	slider := widget.NewSlider(0, 3)
	slider.SetValue(0)
	slider.Step = 1

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(500, 10))

	txt := widget.NewLabel(strconv.Itoa(3 + int(slider.Value)*2))

	sliderContainer := container.NewVBox(
		widget.NewLabel("Mask Size"),
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
		maskSize := 3 + int(value)*2
		txt.SetText(strconv.Itoa(maskSize))
		global.MaskSize = maskSize
	}

	return container.NewPadded(
		container.NewCenter(
			container.NewVBox(sliderBox),
		),
	)
}
