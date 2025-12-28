package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Input(rangeValue uint16, title string, onChange func(string)) *fyne.Container {
	text := canvas.NewText(title, color.White)

	input := widget.NewEntry()
	input.SetText(strconv.Itoa(int(rangeValue)))
	inputWrap := container.NewGridWrap(
		fyne.NewSize(50, input.MinSize().Height),
		input,
	)

	input.OnChanged = func(s string) {
		if onChange != nil {
			onChange(s)
		}
	}

	row := container.NewBorder(nil, nil, text, inputWrap)
	return row
}
