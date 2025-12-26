package core

import (
	"errors"
	"fmt"

	"github.com/sstallion/go-hid"
)

const (
	USB_VID      = 1115
	WHEEL_PID_FS = 22999

	REPORT_LEN = 65

	REPORT_GENERIC_INPUT_OUTPUT      = 0xA3
	REPORT_EFFECT_SETTINGS_FEATURE   = 0x22
	REPORT_HARDWARE_SETTINGS_FEATURE = 0x21
	REPORT_GPIO_SETTINGS_FEATURE     = 0xA1
	REPORT_ADC_SETTINGS_FEATURE      = 0xA2

	DATA_COMMAND_REBOOT        = 0x01
	DATA_COMMAND_SAVE_SETTINGS = 0x02
	DATA_COMMAND_DFU_MODE      = 0x03
	DATA_OVERRIDE_DATA         = 0x10
	DATA_SETTINGS_FIELD_DATA   = 0x14
	DATA_COMMAND_RESET_CENTER  = 0x04
)

type EffectSettings struct {
	MotionRange               uint16
	StaticDampeningStrength   uint16
	SoftStopDampeningStrength uint16
	TotalEffectStrength       byte
	IntegratedSpringStrength  byte
	SoftStopRange             byte
	SoftStopStrength          byte
	DirectXConstantDirection  int8
	DirectXSpringStrength     byte
	DirectXConstantStrength   byte
	DirectXPeriodicStrength   byte
	_padding                  [50]byte
}

type HardwareSettings struct {
	EncoderCPR           uint16
	IntegralGain         uint16
	ProportionalGain     byte
	ForceEnabled         byte
	DebugTorque          byte
	AmplifierGain        byte
	CalibrationMagnitude byte
	CalibrationSpeed     byte
	PowerLimit           byte
	BrakingLimit         byte
	PositionSmoothing    byte
	SpeedBufferSize      byte
	EncoderDirection     int8
	ForceDirection       int8
	PolePairs            byte
	_padding             [47]byte
}

type AdcExtensionSettings struct {
	RAxisMin          [3]uint16
	RAxisMax          [3]uint16
	RAxisSmoothing    [3]byte
	RAxisToButtonLow  [3]byte
	RAxisToButtonHigh [3]byte
	RAxisInvert       [3]byte
	_padding          [40]byte
}

type GpioExtensionSettings struct {
	ExtensionMode     uint8
	PinMode           [10]uint8
	ButtonMode        [32]uint8
	SpiMode           uint8
	SpiLatchMode      uint8
	SpiLatchDelay     uint8
	SpiClkPulseLength uint8
	_padding          [17]byte
}

type FirmwareVersion struct {
	ReleaseType  byte
	ReleaseMajor byte
	ReleaseMinor byte
	ReleasePatch byte
}

type DeviceState struct {
	FirmwareVersion FirmwareVersion
	IsRegistered    byte
	Position        int16
	Torque          int16
	_padding        [55]byte
}

type DirectControl struct {
	SpringForce   int16
	ConstantForce int16
	PeriodicForce int16
	ForceDrop     uint8
}

type Wheel struct {
	dev *hid.Device
}

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
		InstallSudoUdev()
		return err
	}

	w.dev = device
	fmt.Println("Connected!")
	return nil
}

func (w *Wheel) sendCommand(command byte) int {
	fmt.Println(w.dev)
	if w.dev == nil {
		fmt.Println("device not connected")
		return 0
	}

	const reportSize = 65
	report := make([]byte, reportSize)

	report[0] = REPORT_GENERIC_INPUT_OUTPUT
	report[1] = command

	n, err := w.dev.Write(report)
	if err != nil {
		return 0
	}

	return n
}

func (w *Wheel) ResetCenter() int {
	return w.sendCommand(DATA_COMMAND_RESET_CENTER)
}

func (w *Wheel) RebootController() int {
	return w.sendCommand(DATA_COMMAND_REBOOT)
}

func (w *Wheel) SaveAndReboot() int {
	return w.sendCommand(DATA_COMMAND_SAVE_SETTINGS)
}

func (w *Wheel) SwitchToDFU() int {
	return w.sendCommand(DATA_COMMAND_DFU_MODE)
}
