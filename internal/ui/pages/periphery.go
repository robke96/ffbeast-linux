package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

func PeripheryPage(dev *device.Device) *fyne.Container {
	gpioData := dev.Wheel.ReadGPIOSettings()
	resetCenterZ0Value := gpioData.PinMode[8]

	checkbPinZ0 := components.CheckBox(
		"Enable reset center button on pin Z0 (Require reboot)",
		resetCenterZ0Value == 7,
		func(b bool) {
			var val int8
			if b {
				val = 1
			} else {
				val = 0
			}

			dev.Wheel.SetResetCenterOnZ0(val)
		},
	)

	pageContainer := container.NewVBox(
		checkbPinZ0,
	)
	return pageContainer
}
