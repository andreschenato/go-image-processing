package popups

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SavedPopup(w fyne.Window) {
	popup := widget.NewPopUp(
		container.NewCenter(
			canvas.NewText("Image saved successfully", color.White),
		),
		w.Canvas(),
	)

	popup.Resize(fyne.NewSize(popup.MinSize().Width+40, popup.MinSize().Height+20))
	popup.Move(fyne.NewPos(
		w.Canvas().Size().Width-popup.MinSize().Width,
		w.Canvas().Size().Height-popup.MinSize().Height,
	))
	popup.Show()
	time.Sleep(4 * time.Second)
	popup.Hide()
}