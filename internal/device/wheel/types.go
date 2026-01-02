package wheel

import "github.com/sstallion/go-hid"

const (
	USB_VID      = 1115
	WHEEL_PID_FS = 22999

	REPORT_LEN = 65

	REPORT_GENERIC_INPUT_OUTPUT      = 0xA3
	REPORT_EFFECT_SETTINGS_FEATURE   = 0x22
	REPORT_HARDWARE_SETTINGS_FEATURE = 0x21
	REPORT_GPIO_SETTINGS_FEATURE     = 0xA1
	REPORT_ADC_SETTINGS_FEATURE      = 0xA2
	REPORT_FIRMWARE_LICENCE_FEATURE  = 0x25
	REPORT_DEVICE_GAIN               = 0x1D

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
	Padding                   [50]byte
}

type HardwareSettings struct {
	EncoderCPR           uint16
	IntegralGain         uint16
	ProportionalGain     byte
	ForceEnabled         byte // acts as boolean
	DebugTorque          byte // acts as boolean
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
	Padding              [47]byte
}

type AdcExtensionSettings struct {
	RAxisMin          [3]uint16
	RAxisMax          [3]uint16
	RAxisSmoothing    [3]byte
	RAxisToButtonLow  [3]byte
	RAxisToButtonHigh [3]byte
	RAxisInvert       [3]byte
	Padding           [40]byte
}

type GpioExtensionSettings struct {
	ExtensionMode     uint8
	PinMode           [10]uint8
	ButtonMode        [32]uint8
	SpiMode           uint8
	SpiLatchMode      uint8
	SpiLatchDelay     uint8
	SpiClkPulseLength uint8
	Padding           [17]byte
}

type FirmwareVersion struct {
	ReleaseType  byte
	ReleaseMajor byte
	ReleaseMinor byte
	ReleasePatch byte
}

type FirmwareLicence struct {
	FirmwareVersion FirmwareVersion
	SerialKey       [3]uint32
	DeviceId        [3]uint32
	IsRegistered    byte // 0 or 1 acts as boolean flag
	Padding         [35]byte
}

type DeviceState struct {
	FirmwareVersion FirmwareVersion
	IsRegistered    byte // 0 or 1 acts as boolean flag
	Position        int16
	Torque          int16
	Padding         [55]byte
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

type SettingsField uint8

const (
	SETTINGS_FIELD_DIRECT_X_CONSTANT_DIRECTION  SettingsField = 0
	SETTINGS_FIELD_DIRECT_X_SPRING_STRENGTH     SettingsField = 1
	SETTINGS_FIELD_DIRECT_X_CONSTANT_STRENGTH   SettingsField = 2
	SETTINGS_FIELD_DIRECT_X_PERIODIC_STRENGTH   SettingsField = 3
	SETTINGS_FIELD_TOTAL_EFFECT_STRENGTH        SettingsField = 4
	SETTINGS_FIELD_MOTION_RANGE                 SettingsField = 5
	SETTINGS_FIELD_SOFT_STOP_STRENGTH           SettingsField = 6
	SETTINGS_FIELD_SOFT_STOP_RANGE              SettingsField = 7
	SETTINGS_FIELD_STATIC_DAMPENING_STRENGTH    SettingsField = 8
	SETTINGS_FIELD_SOFT_STOP_DAMPENING_STRENGTH SettingsField = 9

	SETTINGS_FIELD_FORCE_ENABLED  SettingsField = 11
	SETTINGS_FIELD_DEBUG_TORQUE   SettingsField = 12
	SETTINGS_FIELD_AMPLIFIER_GAIN SettingsField = 13

	SETTINGS_FIELD_CALIBRATION_MAGNITUDE SettingsField = 15
	SETTINGS_FIELD_CALIBRATION_SPEED     SettingsField = 16
	SETTINGS_FIELD_POWER_LIMIT           SettingsField = 17
	SETTINGS_FIELD_BRAKING_LIMIT         SettingsField = 18
	SETTINGS_FIELD_POSITION_SMOOTHING    SettingsField = 19
	SETTINGS_FIELD_SPEED_BUFFER_SIZE     SettingsField = 20
	SETTINGS_FIELD_ENCODER_DIRECTION     SettingsField = 21
	SETTINGS_FIELD_FORCE_DIRECTION       SettingsField = 22
	SETTINGS_FIELD_POLE_PAIRS            SettingsField = 23
	SETTINGS_FIELD_ENCODER_CPR           SettingsField = 24
	SETTINGS_FIELD_P_GAIN                SettingsField = 25
	SETTINGS_FIELD_I_GAIN                SettingsField = 26

	SETTINGS_FIELD_EXTENSION_MODE       SettingsField = 27
	SETTINGS_FIELD_PIN_MODE             SettingsField = 28
	SETTINGS_FIELD_BUTTON_MODE          SettingsField = 29
	SETTINGS_FIELD_SPI_MODE             SettingsField = 30
	SETTINGS_FIELD_SPI_LATCH_MODE       SettingsField = 31
	SETTINGS_FIELD_SPI_LATCH_DELAY      SettingsField = 32
	SETTINGS_FIELD_SPI_CLK_PULSE_LENGTH SettingsField = 33

	SETTINGS_FIELD_ADC_MIN_DEAD_ZONE  SettingsField = 34
	SETTINGS_FIELD_ADC_MAX_DEAD_ZONE  SettingsField = 35
	SETTINGS_FIELD_ADC_TO_BUTTON_LOW  SettingsField = 36
	SETTINGS_FIELD_ADC_TO_BUTTON_HIGH SettingsField = 37
	SETTINGS_FIELD_ADC_SMOOTHING      SettingsField = 38
	SETTINGS_FIELD_ADC_INVERT         SettingsField = 39

	SETTINGS_FIELD_RESET_CENTER_ON_Z0         SettingsField = 41
	SETTINGS_FIELD_INTEGRATED_SPRING_STRENGTH SettingsField = 43
)
