package generics

import (
	"image_processing/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MaskFormat() *fyne.Container {
	radio := widget.NewRadioGroup([]string{"Square", "Diamond"}, func(val string) {
		switch (val) {
		case "Square":
			global.UseDiamondMask = false
		case "Diamond":
			global.UseDiamondMask = true
		}
	})

	radio.Selected = "Square"
	radio.Horizontal = true

	return container.NewCenter(radio)
}
