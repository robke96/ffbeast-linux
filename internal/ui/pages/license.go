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

func LicensePage(dev *device.Device, onLicenseUpdated func()) *fyne.Container {
	info := widget.NewLabel("")
	status := widget.NewLabel("")

	deviceInput := widget.NewEntry()
	deviceInput.Disable()

	deviceBox := container.NewBorder(
		nil, nil,
		canvas.NewText("Device ID", color.White),
		nil,
		deviceInput,
	)

	var isRegistered bool

	actionBtn := widget.NewButton("Activate", nil)

	refresh := func(message string) {
		if dev == nil || dev.Wheel == nil {
			deviceInput.SetText("")
			info.SetText("Device not connected.")
			actionBtn.SetText("Activate")
			actionBtn.Disable()
			status.SetText(message)
			return
		}

		licenceData := dev.Wheel.ReadFirmwareLicence()
		if licenceData == nil {
			deviceInput.SetText("")
			info.SetText("Failed to read license data from device.")
			actionBtn.SetText("Activate")
			actionBtn.Disable()
			status.SetText(message)
			return
		}

		deviceInput.SetText(arrayToHex(licenceData.DeviceId))
		isRegistered = licenceData.IsRegistered != 0
		if isRegistered {
			info.SetText("PRO license is active on this device.")
			actionBtn.SetText("Deactivate")
		} else {
			info.SetText("Device is not PRO licensed.")
			actionBtn.SetText("Activate")
		}
		actionBtn.Enable()
		status.SetText(message)

		if onLicenseUpdated != nil {
			onLicenseUpdated()
		}
	}

	actionBtn.OnTapped = func() {
		if dev == nil || dev.Wheel == nil {
			refresh("Device not connected.")
			return
		}

		wasRegistered := isRegistered
		actionBtn.Disable()
		status.SetText("Applying license change...")

		var err error
		if wasRegistered {
			err = dev.Wheel.DeactivateLicence()
		} else {
			err = dev.Wheel.ActivateLicence()
		}
		if err != nil {
			refresh("License operation failed: " + err.Error())
			return
		}

		if wasRegistered {
			refresh("License deactivated.")
			return
		}
		refresh("License activated.")
	}

	refresh("")

	page := container.NewVBox(
		info,
		deviceBox,
		actionBtn,
		status,
	)
	return page
}
