package components

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Slider(rangeValue uint16, title string, maxSlider float64, onChange func(float64)) *fyne.Container {
	text := canvas.NewText(title, color.White)

	slider := widget.NewSlider(0, maxSlider)
	slider.SetValue(float64(rangeValue))

	input := widget.NewEntry()
	input.SetText(strconv.Itoa(int(rangeValue)))
	inputWrap := container.NewGridWrap(
		fyne.NewSize(50, input.MinSize().Height),
		input,
	)

	slider.OnChanged = func(v float64) {
		rangeValue = uint16(v)
		input.SetText(fmt.Sprint(rangeValue))
	}

	input.OnChanged = func(s string) {
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			slider.SetValue(v)
		}
	}

	slider.OnChangeEnded = func(v float64) {
		if onChange != nil {
			onChange(v)
		}
	}

	row := container.NewBorder(nil, nil, text, inputWrap, slider)
	return row
}
