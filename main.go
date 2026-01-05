package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui"
)

type forcedVariant struct {
	fyne.Theme
	variant fyne.ThemeVariant
}

func main() {
	a := app.New()
	// force dark theme for now, in future todo: rely on system default?
	a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})

	w := a.NewWindow("FFBeastLinux")

	dev := device.NewDevice()
	ui.NewUI(w, dev)

	w.ShowAndRun()
}
