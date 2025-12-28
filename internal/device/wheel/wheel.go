package wheel

import (
	"errors"
	"fmt"

	"github.com/robke96/ffbeast-linux/internal/device/udev"
	"github.com/sstallion/go-hid"
)

func NewWheel() *Wheel { return &Wheel{} }

func (w *Wheel) Connect() error {
	var path string

	err := hid.Enumerate(USB_VID, WHEEL_PID_FS, func(info *hid.DeviceInfo) error {
		if info.InterfaceNbr == 0 {
			path = info.Path
			return hid.ErrTimeout
		}
		return nil
	})

	if err != nil && err != hid.ErrTimeout {
		return err
	}

	if path == "" {
		return errors.New("Path empty, no device found")
	}

	device, err := hid.OpenPath(path)
	if err != nil {
		udev.InstallSudoUdev()
		return err
	}

	w.dev = device
	fmt.Println("Connected!")
	return nil
}
