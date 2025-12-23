package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/pkg/wheel"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.Resize(fyne.NewSize(300, 400))

	// connect to device
	myWheel := wheel.New()
	myWheel.Connect()
	fmt.Println(myWheel)

	w.SetContent(widget.NewLabel("FFBEAST LINUX :)"))

	// reset center button
	resetCenterBtn := widget.NewButton("Reset Center", func() {
		myWheel.ResetCenter()
	})
	saveAndRebootBtn := widget.NewButton("Save and reboot", func() {
		myWheel.SaveAndReboot()
	})
	switchToDFU := widget.NewButton("Switch to DFU mode", func() {
		myWheel.SwitchToDFU()
	})

	bottomButtons := container.NewHBox(
		saveAndRebootBtn,
		switchToDFU,
		resetCenterBtn,
	)

	content := container.NewBorder(nil, bottomButtons, nil, nil,
		widget.NewLabel("Hello world!"),
	)

	w.SetContent(content)
	w.ShowAndRun()
}
