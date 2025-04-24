package buttons

import (
	"image_processing/global"
	"image_processing/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UpsidedownButton() *fyne.Container {
	btn := widget.NewButton("UpsideDown", func() {
		if *global.ImageOne != nil || *global.ImageTwo != nil {
			global.FinalImage.Image = *service.UpsidedownImage()
			global.FinalImage.Refresh()
		}
	})

	btn.Resize(fyne.NewSize(btn.MinSize().Width+10, btn.MinSize().Height+5))

	return container.NewCenter(btn)
}
