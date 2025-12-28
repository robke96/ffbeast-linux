package wheel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

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

func (w *Wheel) SendSettingReport(field, index int, value any, valueType string) (int, error) {
	if w.dev == nil {
		return 0, errors.New("device not connected")
	}

	buffer := make([]byte, 65)
	buffer[0] = REPORT_GENERIC_INPUT_OUTPUT
	buffer[1] = DATA_SETTINGS_FIELD_DATA
	buffer[2] = byte(field)
	buffer[3] = byte(index)

	dataBuf := bytes.NewBuffer(buffer[4:4])

	switch valueType {
	case "int8_t":
		b := value.(int8)
		dataBuf.WriteByte(byte(b))
	case "uint8_t":
		b := value.(uint8)
		dataBuf.WriteByte(b)
	case "int16_t":
		b := value.(int16)
		binary.Write(dataBuf, binary.LittleEndian, b)
	case "uint16_t":
		b := value.(uint16)
		binary.Write(dataBuf, binary.LittleEndian, b)
	case "float":
		b := value.(float32)
		binary.Write(dataBuf, binary.LittleEndian, b)
	default:
		return 0, errors.New("unsupported type for settings report")
	}

	// copy back to buffer
	copy(buffer[4:], dataBuf.Bytes())

	// send report
	n, err := w.dev.SendOutputReport(buffer)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// buttons
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

// write setting
func (w *Wheel) SetRotationRange(deg uint16) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_MOTION_RANGE),
		0,
		deg,
		"uint16_t",
	)
	return err
}

func (w *Wheel) SetTotalEffectStrength(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_TOTAL_EFFECT_STRENGTH),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetStaticDampening(value uint16) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_STATIC_DAMPENING_STRENGTH),
		0,
		value,
		"uint16_t",
	)
	return err
}

func (w *Wheel) SetDirectXConstantDirection(value int8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_DIRECT_X_CONSTANT_DIRECTION),
		0,
		value,
		"int8_t",
	)
	return err
}

func (w *Wheel) SetResetCenterOnZ0(value int8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_RESET_CENTER_ON_Z0),
		0,
		value,
		"int8_t",
	)
	return err
}

func (w *Wheel) SetEncoderCPR(value uint16) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_ENCODER_CPR),
		0,
		value,
		"uint16_t",
	)
	return err
}

func (w *Wheel) SetPolePairs(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_POLE_PAIRS),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetPGain(value uint8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_P_GAIN),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetIGain(value uint16) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_I_GAIN),
		0,
		value,
		"uint16_t",
	)
	return err
}

func (w *Wheel) SetPowerLimit(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_POWER_LIMIT),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetCalibrationMagnitude(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_CALIBRATION_MAGNITUDE),
		0,
		value,
		"uint8_t",
	)
	return err
}
func (w *Wheel) SetCalibrationSpeed(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_CALIBRATION_SPEED),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetBrakingLimit(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_BRAKING_LIMIT),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetEnableForces(value byte) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_FORCE_ENABLED),
		0,
		value,
		"uint8_t",
	)
	return err
}

func (w *Wheel) SetInvertJoystickOutput(value int8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_ENCODER_DIRECTION),
		0,
		value,
		"int8_t",
	)
	return err
}

func (w *Wheel) SetInvertForceOutput(value int8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_FORCE_DIRECTION),
		0,
		value,
		"int8_t",
	)
	return err
}

func (w *Wheel) SetDebugForces(value uint8) error {
	_, err := w.SendSettingReport(
		int(SETTINGS_FIELD_DEBUG_TORQUE),
		0,
		value,
		"uint8_t",
	)
	return err
}
