package pages

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return -1
}

func ControllerPage(dev *device.Device) *fyne.Container {
	controlData := dev.Wheel.ReadHardwareSettings()

	checkEnableForces := components.CheckBox(
		"Enable forces (Require reboot)",
		controlData.ForceEnabled == 1,
		func(b bool) {
			dev.Wheel.SetEnableForces(BoolToByte(b))
		},
	)

	checkInvertJoy := components.CheckBox(
		"Invert joystick output",
		controlData.EncoderDirection == 1,
		func(b bool) {
			dev.Wheel.SetInvertJoystickOutput(BoolToInt8(b))
		},
	)

	checkInvertForce := components.CheckBox(
		"Invert force output",
		controlData.ForceDirection == 1,
		func(b bool) {
			dev.Wheel.SetInvertForceOutput(BoolToInt8(b))
		},
	)

	checkDebugForces := components.CheckBox(
		"Debug forces",
		controlData.DebugTorque == 1,
		func(b bool) {
			dev.Wheel.SetDebugForces(BoolToByte(b))
		},
	)

	encodercpr := components.Input(
		controlData.EncoderCPR,
		"Encoder CPR",
		func(s string) {
			num, err := strconv.ParseUint(s, 10, 16)
			if err != nil {
				return
			}
			dev.Wheel.SetEncoderCPR(uint16(num))
		},
	)

	polepairs := components.Input(
		uint16(controlData.PolePairs),
		"Pole pairs",
		func(s string) {
			num, err := strconv.ParseUint(s, 10, 8)
			if err != nil {
				return
			}
			dev.Wheel.SetPolePairs(byte(num))
		},
	)

	pgain := components.Slider(
		uint16(controlData.ProportionalGain),
		"P gain",
		100,
		func(f float64) {
			dev.Wheel.SetPGain(byte(f))
		},
	)

	igain := components.Slider(
		uint16(controlData.ProportionalGain),
		"I gain",
		500,
		func(f float64) {
			dev.Wheel.SetIGain(uint16(f))
		},
	)

	powerlimit := components.Slider(
		uint16(controlData.PowerLimit),
		"Power limit (%)",
		100,
		func(f float64) {
			dev.Wheel.SetPowerLimit(byte(f))
		},
	)

	calMagnitude := components.Slider(
		uint16(controlData.CalibrationMagnitude),
		"Calibration magnitude (%)",
		100,
		func(f float64) {
			dev.Wheel.SetCalibrationMagnitude(byte(f))
		},
	)

	calSpeed := components.Slider(
		uint16(controlData.CalibrationSpeed),
		"Calibration speed (%)",
		100,
		func(f float64) {
			dev.Wheel.SetCalibrationSpeed(byte(f))
		},
	)

	brakingLimit := components.Slider(
		uint16(controlData.BrakingLimit),
		"Braking resistor limit (%)",
		100,
		func(f float64) {
			dev.Wheel.SetBrakingLimit(byte(f))
		},
	)

	pageContainer := container.NewVBox(
		checkEnableForces,
		checkInvertJoy,
		checkInvertForce,
		checkDebugForces,
		encodercpr,
		polepairs,
		pgain,
		igain,
		powerlimit,
		calMagnitude,
		calSpeed,
		brakingLimit,
	)
	return pageContainer
}
