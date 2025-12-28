package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("FFBeastLinux")

	dev := device.NewDevice()
	ui.NewUI(w, dev)

	w.ShowAndRun()
}
