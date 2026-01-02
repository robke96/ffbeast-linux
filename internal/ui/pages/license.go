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

func LicensePage(dev *device.Device) *fyne.Container {
	info := widget.NewLabel("This app is only for basic version at this moment.")

	licenceData := dev.Wheel.ReadFirmwareLicence()

	// TO-DO: not reading, int -> hex
	deviceId := fmt.Sprintf("%x", licenceData.DeviceId)

	deviceInput := widget.NewEntry()
	deviceInput.SetText(deviceId)
	deviceInput.Disable()

	deviceBox := container.NewBorder(
		nil, nil,
		canvas.NewText("Device ID", color.White),
		nil,
		deviceInput,
	)

	serialInput := widget.NewEntry()
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
