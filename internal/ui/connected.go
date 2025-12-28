package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/pages"
)

func bottomButtons(dev *device.Device) fyne.CanvasObject {
	resetWheelCenterBtn := widget.NewButton("Reset wheel center", func() {
		dev.Wheel.ResetCenter()
	})

	saveRebootBtn := widget.NewButton("Save and reboot", func() {
		dev.Wheel.SaveAndReboot()
	})

	switchDFUmodeBtn := widget.NewButton("Switch to DFU mode", func() {
		dev.Wheel.SwitchToDFU()
	})

	rebootControllerBtn := widget.NewButton("Reboot controller", func() {
		dev.Wheel.RebootController()
	})

	column := container.NewVBox(
		resetWheelCenterBtn,
		saveRebootBtn,
		switchDFUmodeBtn,
		rebootControllerBtn,
	)
	return column
}

func ConnectedPage(dev *device.Device) *fyne.Container {
	tabs := container.NewAppTabs(
		container.NewTabItem("Effects", pages.EffectsPage(dev)),
		container.NewTabItem("Periphery", pages.PeripheryPage(dev)),
		container.NewTabItem("Controller", pages.ControllerPage(dev)),
		container.NewTabItem("License", pages.LicensePage()),
	)

	content := container.NewBorder(
		nil,
		bottomButtons(dev),
		nil,
		nil,
		tabs,
	)

	return content
}
