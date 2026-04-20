package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/device/wheel"
)

const BaseWindowTitle = "FFBeastLinux"

func buildWindowTitle(licence *wheel.FirmwareLicence) string {
	if licence == nil {
		return BaseWindowTitle
	}

	title := BaseWindowTitle
	if licence.IsRegistered != 0 {
		title += " PRO"
	}

	version := licence.FirmwareVersion
	return fmt.Sprintf("%s v%d.%d.%d", title, version.ReleaseMajor, version.ReleaseMinor, version.ReleasePatch)
}

func setWindowTitle(w fyne.Window, dev *device.Device) {
	if dev == nil || !dev.Connected || dev.Wheel == nil {
		w.SetTitle(BaseWindowTitle)
		return
	}

	licenceData := dev.Wheel.ReadFirmwareLicence()
	w.SetTitle(buildWindowTitle(licenceData))
}
