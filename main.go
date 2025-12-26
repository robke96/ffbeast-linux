package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/robke96/ffbeast-linux/internal/core"
	"github.com/robke96/ffbeast-linux/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("FFBeastLinux")

	appState := core.NewAppState()
	ui.NewUI(w, appState)

	w.ShowAndRun()
}
