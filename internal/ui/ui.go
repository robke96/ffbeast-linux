package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/device/wheel"
)

func NewUI(w fyne.Window, dev *device.Device) {
	w.Resize(fyne.NewSize(600, 800))
	wheelIcon := NewWheelIcon("assets/wheel.webp", 28)

	// first init when app starts up
	myWheel := wheel.NewWheel()
	err := myWheel.Connect()

	if err == nil {
		dev.Connected = true
		dev.Wheel = myWheel
		if effect := dev.Wheel.ReadEffectSettings(); effect != nil {
			wheelIcon.SetMotionRange(effect.MotionRange)
		}
		setWindowTitle(w, dev)
		w.SetContent(ConnectedPage(dev, func() {
			setWindowTitle(w, dev)
		}, wheelIcon.Object()))
	} else {
		w.SetTitle(BaseWindowTitle)
		w.SetContent(WaitingPage())
	}

	// auto reconnect ping logic
	go func() {
		missedStateReads := 0
		for {
			if !dev.Connected {
				err := myWheel.Connect()

				if err == nil {
					dev.Connected = true
					dev.Wheel = myWheel
					missedStateReads = 0

					fyne.Do(func() {
						if effect := dev.Wheel.ReadEffectSettings(); effect != nil {
							wheelIcon.SetMotionRange(effect.MotionRange)
						}
						setWindowTitle(w, dev)
						w.SetContent(ConnectedPage(dev, func() {
							setWindowTitle(w, dev)
						}, wheelIcon.Object()))
					})
				}
				time.Sleep(1 * time.Second)
			} else {
				state := myWheel.ReadState()
				if state != nil {
					missedStateReads = 0
					wheelIcon.SetPosition(state.Position)
					time.Sleep(5 * time.Millisecond)
					continue
				}

				missedStateReads++
				if missedStateReads < 8 {
					time.Sleep(5 * time.Millisecond)
					continue
				}

				if !myWheel.IsConnected() {
					dev.Connected = false
					dev.Wheel = nil
					missedStateReads = 0
					fyne.Do(func() {
						w.SetTitle(BaseWindowTitle)
						w.SetContent(WaitingPage())
					})
				} else {
					missedStateReads = 0
				}
			}
		}
	}()
}
