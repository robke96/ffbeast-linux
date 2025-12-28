package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

func PeripheryPage(dev *device.Device) *fyne.Container {
	checkbPinZ0 := components.CheckBox(
		"Enable reset center button on pin Z0 (Require reboot)",
		false,
		func(b bool) {
			var val int8
			if b {
				val = 1
			} else {
				val = -1
			}

			dev.Wheel.SetResetCenterOnZ0(val)
		},
	)
	checkbPinZ0.Disable()

	pageContainer := container.NewVBox(
		widget.NewLabel("CURRENTLY NOT WORKING"),
		checkbPinZ0,
	)
	return pageContainer
}
