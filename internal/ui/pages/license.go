package pages

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
)

func arrayToHex(arr [3]uint32) string {
	return fmt.Sprintf("%08X-%08X-%08X", arr[0], arr[1], arr[2])
}

func LicensePage(dev *device.Device) *fyne.Container {
	info := widget.NewLabel("This app is only for basic version at this moment.")
	licenceData := dev.Wheel.ReadFirmwareLicence()

	deviceId := arrayToHex(licenceData.DeviceId)
	deviceInput := widget.NewEntry()
	deviceInput.SetText(deviceId)
	deviceInput.Disable()

	deviceBox := container.NewBorder(
		nil, nil,
		canvas.NewText("Device ID", color.White),
		nil,
		deviceInput,
	)

	serialId := arrayToHex(licenceData.SerialKey)
	serialInput := widget.NewEntry()
	serialInput.SetText(serialId)

	serialBox := container.NewBorder(
		nil, nil,
		canvas.NewText("Serial Key", color.White),
		nil,
		serialInput,
	)

	activateBtn := widget.NewButton("Activate", func() {
		// for now nothing
	})
	activateBtn.Disable()

	page := container.NewVBox(
		info,
		deviceBox,
		serialBox,
		activateBtn,
	)
	return page
}
