package ui

import (
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"github.com/avast/retry-go"
	"github.com/robke96/ffbeast-linux/internal/core"
)

func NewUI(w fyne.Window, state *core.AppState) {
	w.Resize(fyne.NewSize(600, 800))

	myWheel := core.NewWheel()
	err := myWheel.Connect()
	if err == nil {
		state.DeviceConnected = true
		state.Wheel = myWheel
		w.SetContent(ConnectedPage(state))
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
			state.DeviceConnected = true
			state.Wheel = myWheel
			state.CurrentPage = "effects"

			fyne.Do(func() {
				w.Canvas().Refresh(ConnectedPage(state))
				w.SetContent(ConnectedPage(state))
			})
		} else {
			fmt.Println("failed")
		}
	}()
}
