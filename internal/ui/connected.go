package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/robke96/ffbeast-linux/internal/core"
)

func ConnectedPage(state *core.AppState) *fyne.Container {
	tabs := container.NewAppTabs(
		container.NewTabItem("Effects", EffectsPage()),
		container.NewTabItem("Periphery", PeripheryPage()),
		container.NewTabItem("Controller", ControllerPage()),
		container.NewTabItem("License", LicensePage()),
	)

	return container.NewStack(tabs)
}
