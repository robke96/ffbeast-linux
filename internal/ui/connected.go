package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/core"
)

func bottomButtons(state *core.AppState) fyne.CanvasObject {
	resetWheelCenterBtn := widget.NewButton("Reset wheel center", func() {
		state.Wheel.ResetCenter()
	})

	saveRebootBtn := widget.NewButton("Save and reboot", func() {
		state.Wheel.SaveAndReboot()
	})

	switchDFUmodeBtn := widget.NewButton("Switch to DFU mode", func() {
		state.Wheel.SwitchToDFU()
	})

	rebootControllerBtn := widget.NewButton("Reboot controller", func() {
		state.Wheel.RebootController()
	})

	column := container.NewVBox(
		resetWheelCenterBtn,
		saveRebootBtn,
		switchDFUmodeBtn,
		rebootControllerBtn,
	)
	return column
}

func ConnectedPage(state *core.AppState) *fyne.Container {
	tabs := container.NewAppTabs(
		container.NewTabItem("Effects", EffectsPage()),
		container.NewTabItem("Periphery", PeripheryPage()),
		container.NewTabItem("Controller", ControllerPage()),
		container.NewTabItem("License", LicensePage()),
	)

	content := container.NewBorder(
		nil,
		bottomButtons(state),
		nil,
		nil,
		tabs,
	)

	return content
}
