package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/device/wheel"
)

func NewUI(w fyne.Window, dev *device.Device) {
	w.Resize(fyne.NewSize(600, 800))

	// first init when app starts up
	myWheel := wheel.NewWheel()
	err := myWheel.Connect()

	if err == nil {
		dev.Connected = true
		dev.Wheel = myWheel
		setWindowTitle(w, dev)
		w.SetContent(ConnectedPage(dev, func() {
			setWindowTitle(w, dev)
		}))
	} else {
		w.SetTitle(BaseWindowTitle)
		w.SetContent(WaitingPage())
	}

	// auto reconnect ping logic
	go func() {
		for {
			if !dev.Connected {
				err := myWheel.Connect()

				if err == nil {
					dev.Connected = true
					dev.Wheel = myWheel

					fyne.Do(func() {
						setWindowTitle(w, dev)
						w.SetContent(ConnectedPage(dev, func() {
							setWindowTitle(w, dev)
						}))
					})
				}
			} else {
				if !myWheel.IsConnected() {
					dev.Connected = false
					dev.Wheel = nil
					fyne.Do(func() {
						w.SetTitle(BaseWindowTitle)
						w.SetContent(WaitingPage())
					})
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()
}
