package sliders

import (
	"image_processing/global"
	"image_processing/service"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BrightSlider() *fyne.Container {
	slider := widget.NewSlider(-100, 100)
	slider.SetValue(0)

	fixedLayout := layout.NewGridWrapLayout(fyne.NewSize(500, 10))

	txt := widget.NewLabel(strconv.Itoa(int(slider.Value)))

	sliderContainer := container.New(fixedLayout, slider, txt)

	slider.OnChanged = func(value float64) {
		txt.SetText(strconv.Itoa(int(value)))
		if *global.ImageOne != nil || *global.ImageTwo != nil {
			global.FinalImage.Image = service.BrightImage(value)
			global.FinalImage.Refresh()
		}
	}

	return container.NewCenter(container.NewHBox(sliderContainer))
}
