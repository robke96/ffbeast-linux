package components

import (
	"fyne.io/fyne/v2/widget"
)

func CheckBox(title string, isChecked bool, changed func(bool)) *widget.Check {
	item := widget.NewCheck(title, func(b bool) {
		if changed != nil {
			changed(b)
		}
	})
	item.SetChecked(isChecked)

	return item
}
