package ui

import (
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"github.com/avast/retry-go"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/device/wheel"
)

func NewUI(w fyne.Window, dev *device.Device) {
	w.Resize(fyne.NewSize(600, 800))

	myWheel := wheel.NewWheel()
	err := myWheel.Connect()
	if err == nil {
		dev.Connected = true
		dev.Wheel = myWheel
		w.SetContent(ConnectedPage(dev))
		return
	}

	w.SetContent(WaitingPage())
	go func() {
		err := retry.Do(func() error {
			err := myWheel.Connect()
			if err != nil {
				fmt.Printf("Connect failed - %s\n", err)
				return errors.New("connection failed")
			}

			return nil
		},
			retry.Attempts(15),
			retry.Delay(3*time.Second),
		)

		if err == nil {
			dev.Connected = true
			dev.Wheel = myWheel

			fyne.Do(func() {
				w.Canvas().Refresh(ConnectedPage(dev))
				w.SetContent(ConnectedPage(dev))
			})
		} else {
			fmt.Println("failed")
		}
	}()
}
