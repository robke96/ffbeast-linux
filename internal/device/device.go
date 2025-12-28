package device

import "github.com/robke96/ffbeast-linux/internal/device/wheel"

type Device struct {
	Connected bool
	Wheel     *wheel.Wheel
}

func NewDevice() *Device {
	return &Device{
		Connected: false,
		Wheel:     nil,
	}
}
