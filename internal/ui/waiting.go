package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func WaitingPage() *fyne.Container {
	text := canvas.NewText("Waiting for device...", color.White)
	text.TextSize = 16

	return container.NewCenter(text)
}
