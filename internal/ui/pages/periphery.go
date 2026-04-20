package pages

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/device/wheel"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

type pinModeOption struct {
	Label string
	Value uint8
}

type collapsibleSection struct {
	title        string
	disabledHint string
	detail       fyne.CanvasObject
	enabled      bool
	isOpen       bool
	header       *widget.Button
	body         *fyne.Container
	hint         *widget.Label
	container    *fyne.Container
}

var spiExtensionModeOptions = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "Custom", Value: 1},
}

var spiModeOptions = []pinModeOption{
	{Label: "Mode 0", Value: 0},
	{Label: "Mode 1", Value: 1},
	{Label: "Mode 2", Value: 2},
	{Label: "Mode 3", Value: 3},
}

var latchModeOptions = []pinModeOption{
	{Label: "Latch UP", Value: 0},
	{Label: "Latch DOWN", Value: 1},
}

var buttonModeOptions = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "Normal", Value: 1},
	{Label: "Inverted", Value: 2},
}

var pinModeOptionsPin123 = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "Generic Button", Value: 1},
	{Label: "Analog", Value: 2},
	{Label: "SPI CS", Value: 3},
	{Label: "SPI SCK", Value: 4},
	{Label: "SPI MISO", Value: 5},
	{Label: "Enable effects switch", Value: 6},
	{Label: "Reset center button", Value: 7},
	{Label: "Effect led", Value: 9},
	{Label: "Reboot", Value: 10},
}

var pinModeOptionsPin4to7 = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "Generic Button", Value: 1},
	{Label: "SPI CS", Value: 3},
	{Label: "SPI SCK", Value: 4},
	{Label: "SPI MISO", Value: 5},
	{Label: "Enable effects switch", Value: 6},
	{Label: "Reset center button", Value: 7},
	{Label: "Effect led", Value: 9},
	{Label: "Reboot", Value: 10},
}

var pinModeOptionsPin8 = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "Generic Button", Value: 1},
	{Label: "SPI CS", Value: 3},
	{Label: "SPI SCK", Value: 4},
	{Label: "SPI MISO", Value: 5},
	{Label: "Enable effects switch", Value: 6},
	{Label: "Reset center button", Value: 7},
	{Label: "Braking PWM", Value: 8},
	{Label: "Effect led", Value: 9},
	{Label: "Reboot", Value: 10},
}

var pinModeOptionsEncoder = []pinModeOption{
	{Label: "Disabled", Value: 0},
	{Label: "SPI CS", Value: 3},
	{Label: "SPI SCK", Value: 4},
	{Label: "SPI MISO", Value: 5},
	{Label: "Enable effects switch", Value: 6},
	{Label: "Reset center button", Value: 7},
	{Label: "Effect led", Value: 9},
	{Label: "Reboot", Value: 10},
}

func pinModeOptionsByIndex(pinIndex uint8) []pinModeOption {
	switch pinIndex {
	case 0, 1, 2:
		return pinModeOptionsPin123
	case 3, 4, 5, 6:
		return pinModeOptionsPin4to7
	case 7:
		return pinModeOptionsPin8
	case 8, 9:
		return pinModeOptionsEncoder
	default:
		return pinModeOptionsPin4to7
	}
}

func pinModeLabels(options []pinModeOption) []string {
	labels := make([]string, 0, len(options))
	for _, option := range options {
		labels = append(labels, option.Label)
	}
	return labels
}

func pinModeLabelByValue(options []pinModeOption, value uint8) string {
	for _, option := range options {
		if option.Value == value {
			return option.Label
		}
	}
	return options[0].Label
}

func pinModeValueByLabel(options []pinModeOption, label string) (uint8, bool) {
	for _, option := range options {
		if option.Label == label {
			return option.Value, true
		}
	}
	return 0, false
}

func selectRow(label string, options []pinModeOption, currentValue uint8, onChange func(uint8)) fyne.CanvasObject {
	selectInput := widget.NewSelect(pinModeLabels(options), nil)
	selectInput.SetSelected(pinModeLabelByValue(options, currentValue))
	selectInput.OnChanged = func(s string) {
		value, ok := pinModeValueByLabel(options, s)
		if !ok {
			return
		}
		onChange(value)
	}

	return container.NewBorder(
		nil,
		nil,
		canvas.NewText(label, color.White),
		nil,
		selectInput,
	)
}

func uint8InputRow(label string, currentValue uint8, onChange func(uint8)) fyne.CanvasObject {
	return components.Input(
		uint16(currentValue),
		label,
		func(s string) {
			n, err := strconv.ParseUint(s, 10, 8)
			if err != nil {
				return
			}
			onChange(uint8(n))
		},
	)
}

func tableUint16Cell(initial uint16, min uint16, max uint16, onChange func(uint16)) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetText(strconv.Itoa(int(initial)))
	input.OnChanged = func(s string) {
		n, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return
		}

		value := uint16(n)
		if value < min || value > max {
			return
		}

		onChange(value)
	}

	return input
}

func tableUint8Cell(initial uint8, max uint8, onChange func(uint8)) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetText(strconv.Itoa(int(initial)))
	input.OnChanged = func(s string) {
		n, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return
		}

		value := uint8(n)
		if value > max {
			return
		}

		onChange(value)
	}

	return input
}

func tableInvertCell(initial bool, onChange func(bool)) fyne.CanvasObject {
	check := widget.NewCheck("", func(b bool) {
		onChange(b)
	})
	check.SetChecked(initial)
	return container.NewCenter(check)
}

func fixedWidthCell(obj fyne.CanvasObject, width float32) fyne.CanvasObject {
	return container.NewGridWrap(
		fyne.NewSize(width, obj.MinSize().Height),
		obj,
	)
}

func newCollapsibleSection(title string, detail fyne.CanvasObject, enabled bool, disabledHint string) *collapsibleSection {
	section := &collapsibleSection{
		title:        title,
		disabledHint: disabledHint,
		detail:       detail,
	}

	section.body = container.NewVBox()
	section.hint = widget.NewLabel("")
	section.header = widget.NewButton("", func() {
		section.setOpen(!section.isOpen)
	})
	section.container = container.NewVBox(section.header, section.body, section.hint)

	section.setEnabled(enabled)
	return section
}

func (s *collapsibleSection) setOpen(open bool) {
	if !s.enabled {
		return
	}

	s.isOpen = open
	if s.isOpen {
		s.header.SetText("Close " + s.title)
		s.body.Objects = []fyne.CanvasObject{s.detail}
		s.body.Refresh()
		s.container.Refresh()
		return
	}

	s.header.SetText("Open " + s.title)
	s.body.Objects = nil
	s.body.Refresh()
	s.container.Refresh()
}

func (s *collapsibleSection) setEnabled(enabled bool) {
	s.enabled = enabled

	if !enabled {
		s.header.Disable()
		s.isOpen = false
		s.header.SetText("Open " + s.title)
		s.body.Objects = nil
		s.body.Refresh()
		s.hint.SetText(s.disabledHint)
		s.container.Refresh()
		return
	}

	s.header.Enable()
	s.hint.SetText("")
	s.setOpen(false)
	s.container.Refresh()
}

func hasAnyButtonPin(pinModes [10]uint8) bool {
	for _, mode := range pinModes {
		if mode == 1 {
			return true
		}
	}
	return false
}

func hasAnyAnalogPin(pinModes [10]uint8) bool {
	for _, mode := range pinModes {
		if mode == 2 {
			return true
		}
	}
	return false
}

func buildProGPIOSection(dev *device.Device, gpioData *wheel.GpioExtensionSettings, onPinModeChanged func(uint8, uint8)) fyne.CanvasObject {
	pinNames := []string{
		"Pin 1",
		"Pin 2",
		"Pin 3",
		"Pin 4",
		"Pin 5",
		"Pin 6",
		"Pin 7",
		"Pin 8",
		"Encoder Z0",
		"Encoder Z1",
	}

	rows := []fyne.CanvasObject{
		widget.NewLabel("GPIO modes:"),
	}
	for i, pinName := range pinNames {
		pinIndex := uint8(i)
		rows = append(
			rows,
			selectRow(
				pinName,
				pinModeOptionsByIndex(pinIndex),
				gpioData.PinMode[i],
				func(v uint8) {
					dev.Wheel.SetPinMode(pinIndex, v)
					if onPinModeChanged != nil {
						onPinModeChanged(pinIndex, v)
					}
				},
			),
		)
	}

	return container.NewVBox(rows...)
}

func buildButtonModesSection(dev *device.Device, gpioData *wheel.GpioExtensionSettings) fyne.CanvasObject {
	rows := []fyne.CanvasObject{
		widget.NewLabel("GPIO buttons:"),
	}

	for i := range 8 {
		buttonIndex := uint8(i)
		rows = append(
			rows,
			selectRow(
				"Button "+strconv.Itoa(i+1),
				buttonModeOptions,
				gpioData.ButtonMode[i],
				func(v uint8) {
					dev.Wheel.SetButtonMode(buttonIndex, v)
				},
			),
		)
	}

	return container.NewVBox(rows...)
}

func buildProSPISection(dev *device.Device, gpioData *wheel.GpioExtensionSettings) fyne.CanvasObject {
	return container.NewVBox(
		selectRow(
			"Spi extension mode",
			spiExtensionModeOptions,
			gpioData.ExtensionMode,
			func(v uint8) { dev.Wheel.SetExtensionMode(v) },
		),
		selectRow(
			"SPI mode",
			spiModeOptions,
			gpioData.SpiMode,
			func(v uint8) { dev.Wheel.SetSpiMode(v) },
		),
		selectRow(
			"Latch mode",
			latchModeOptions,
			gpioData.SpiLatchMode,
			func(v uint8) { dev.Wheel.SetSpiLatchMode(v) },
		),
		uint8InputRow(
			"Latch delay (microseconds)",
			gpioData.SpiLatchDelay,
			func(v uint8) { dev.Wheel.SetSpiLatchDelay(v) },
		),
		uint8InputRow(
			"SCK pulse length (microseconds)",
			gpioData.SpiClkPulseLength,
			func(v uint8) { dev.Wheel.SetSpiClkPulseLength(v) },
		),
	)
}

func buildProAnalogAxesSection(dev *device.Device, adcData *wheel.AdcExtensionSettings) fyne.CanvasObject {
	const (
		axisColWidth  float32 = 30
		valueColWidth float32 = 86
		invertWidth   float32 = 48
	)

	headerRow := container.NewHBox(
		fixedWidthCell(widget.NewLabel(""), axisColWidth),
		fixedWidthCell(widget.NewLabel("Low DZ"), valueColWidth),
		fixedWidthCell(widget.NewLabel("High DZ"), valueColWidth),
		fixedWidthCell(widget.NewLabel("Smooth %"), valueColWidth),
		fixedWidthCell(widget.NewLabel("Btn low %"), valueColWidth),
		fixedWidthCell(widget.NewLabel("Btn high %"), valueColWidth),
		fixedWidthCell(widget.NewLabel("Inv"), invertWidth),
	)

	rows := []fyne.CanvasObject{headerRow}

	axes := []string{"Rx", "Ry", "Rz"}
	for i, axis := range axes {
		axisIndex := uint8(i)

		row := container.NewHBox(
			fixedWidthCell(canvas.NewText(axis, color.White), axisColWidth),
			fixedWidthCell(
				tableUint16Cell(
					adcData.RAxisMin[i],
					0,
					32767,
					func(v uint16) { dev.Wheel.SetAdcMinDeadZone(axisIndex, v) },
				),
				valueColWidth,
			),
			fixedWidthCell(
				tableUint16Cell(
					adcData.RAxisMax[i],
					0,
					32767,
					func(v uint16) { dev.Wheel.SetAdcMaxDeadZone(axisIndex, v) },
				),
				valueColWidth,
			),
			fixedWidthCell(
				tableUint8Cell(
					adcData.RAxisSmoothing[i],
					100,
					func(v uint8) { dev.Wheel.SetAdcSmoothing(axisIndex, v) },
				),
				valueColWidth,
			),
			fixedWidthCell(
				tableUint8Cell(
					adcData.RAxisToButtonLow[i],
					100,
					func(v uint8) { dev.Wheel.SetAdcToButtonLow(axisIndex, v) },
				),
				valueColWidth,
			),
			fixedWidthCell(
				tableUint8Cell(
					adcData.RAxisToButtonHigh[i],
					100,
					func(v uint8) { dev.Wheel.SetAdcToButtonHigh(axisIndex, v) },
				),
				valueColWidth,
			),
			fixedWidthCell(
				tableInvertCell(
					adcData.RAxisInvert[i] == 1,
					func(b bool) {
						var v uint8
						if b {
							v = 1
						} else {
							v = 0
						}
						dev.Wheel.SetAdcInvert(axisIndex, v)
					},
				),
				invertWidth,
			),
		)

		rows = append(rows, row)
	}

	return container.NewVBox(rows...)
}

func PeripheryPage(dev *device.Device) *fyne.Container {
	if dev == nil || dev.Wheel == nil {
		return container.NewVBox(widget.NewLabel("Device not connected."))
	}

	gpioData := dev.Wheel.ReadGPIOSettings()
	if gpioData == nil {
		return container.NewVBox(widget.NewLabel("Failed to read periphery settings from device."))
	}

	licenceData := dev.Wheel.ReadFirmwareLicence()
	isPro := licenceData != nil && licenceData.IsRegistered != 0

	if isPro {
		adcData := dev.Wheel.ReadADCSettings()
		if adcData == nil {
			return container.NewVBox(widget.NewLabel("Failed to read analog axes settings from device."))
		}

		currentPinModes := gpioData.PinMode

		buttonModesSection := newCollapsibleSection(
			"Button modes",
			buildButtonModesSection(dev, gpioData),
			hasAnyButtonPin(currentPinModes),
			"Set any GPIO pin mode to Generic Button to unlock this section.",
		)

		analogAxesSection := newCollapsibleSection(
			"Analog axes",
			buildProAnalogAxesSection(dev, adcData),
			hasAnyAnalogPin(currentPinModes),
			"Set any GPIO pin mode to Analog to unlock this section.",
		)

		gpioSection := newCollapsibleSection(
			"GPIO (Require reboot)",
			buildProGPIOSection(dev, gpioData, func(index, mode uint8) {
				previousButtonState := hasAnyButtonPin(currentPinModes)
				previousAnalogState := hasAnyAnalogPin(currentPinModes)
				currentPinModes[index] = mode
				currentButtonState := hasAnyButtonPin(currentPinModes)
				currentAnalogState := hasAnyAnalogPin(currentPinModes)
				if previousButtonState != currentButtonState {
					buttonModesSection.setEnabled(currentButtonState)
				}
				if previousAnalogState != currentAnalogState {
					analogAxesSection.setEnabled(currentAnalogState)
				}
			}),
			true,
			"",
		)

		spiSection := newCollapsibleSection(
			"SPI extension",
			buildProSPISection(dev, gpioData),
			true,
			"",
		)

		content := container.NewVBox(
			gpioSection.container,
			spiSection.container,
			buttonModesSection.container,
			analogAxesSection.container,
		)
		return container.NewStack(container.NewVScroll(content))
	}

	resetCenterZ0Value := gpioData.PinMode[8]

	checkbPinZ0 := components.CheckBox(
		"Enable reset center button on pin Z0 (Require reboot)",
		resetCenterZ0Value == 7,
		func(b bool) {
			var val int8
			if b {
				val = 1
			} else {
				val = 0
			}

			dev.Wheel.SetResetCenterOnZ0(val)
		},
	)

	pageContainer := container.NewVBox(
		checkbPinZ0,
	)
	return pageContainer
}
