package ui

import "fyne.io/fyne/v2/widget"

func WaitingPage() *widget.Label {
	return widget.NewLabel("Waiting for device..")
}
