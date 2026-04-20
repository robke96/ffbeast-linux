package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

func EffectsPage(dev *device.Device) *fyne.Container {
	if dev == nil || dev.Wheel == nil {
		return container.NewVBox(widget.NewLabel("Device not connected."))
	}

	effectSettings := dev.Wheel.ReadEffectSettings()
	if effectSettings == nil {
		return container.NewVBox(widget.NewLabel("Failed to read effect settings from device."))
	}

	licenceData := dev.Wheel.ReadFirmwareLicence()
	isPro := licenceData != nil && licenceData.IsRegistered != 0

	motionRangeSlider := components.Slider(
		effectSettings.MotionRange,
		"Motion range (degrees)",
		1080,
		func(f float64) {
			dev.Wheel.SetRotationRange(uint16(f))
		},
	)

	totalStrengthSlider := components.Slider(
		uint16(effectSettings.TotalEffectStrength),
		"Total effect strength (%)",
		100,
		func(f float64) {
			dev.Wheel.SetTotalEffectStrength(byte(f))
		},
	)

	staticDampeningSlider := components.Slider(
		effectSettings.StaticDampeningStrength,
		"Static dampening (%)",
		100,
		func(f float64) {
			dev.Wheel.SetStaticDampening(uint16(f))
		},
	)

	commonItems := []fyne.CanvasObject{
		widget.NewLabel("Common"),
		motionRangeSlider,
		totalStrengthSlider,
		staticDampeningSlider,
	}

	if isPro {
		integratedSpringStrengthSlider := components.Slider(
			uint16(effectSettings.IntegratedSpringStrength),
			"Integrated spring strength (%)",
			100,
			func(f float64) {
				dev.Wheel.SetIntegratedSpringStrength(byte(f))
			},
		)

		softStopStrengthSlider := components.Slider(
			uint16(effectSettings.SoftStopStrength),
			"Soft stop strength (%)",
			100,
			func(f float64) {
				dev.Wheel.SetSoftStopStrength(byte(f))
			},
		)

		softStopRangeSlider := components.Slider(
			uint16(effectSettings.SoftStopRange),
			"Soft stop range (degrees)",
			100,
			func(f float64) {
				dev.Wheel.SetSoftStopRange(byte(f))
			},
		)

		softStopDampeningSlider := components.Slider(
			effectSettings.SoftStopDampeningStrength,
			"Soft stop dampening (%)",
			100,
			func(f float64) {
				dev.Wheel.SetSoftStopDampening(uint16(f))
			},
		)

		commonItems = append(
			commonItems,
			integratedSpringStrengthSlider,
			softStopStrengthSlider,
			softStopRangeSlider,
			softStopDampeningSlider,
		)
	}

	commonSettings := container.NewVBox(commonItems...)

	// TO-DO needs optimization, maybe possible to use bool instead of int8?
	constantForceBool := effectSettings.DirectXConstantDirection == 1
	checkbDirectXInvertConstantForce := components.CheckBox(
		"Invert constant force",
		constantForceBool,
		func(b bool) {
			var val int8
			if b {
				val = 1
			} else {
				val = -1
			}

			dev.Wheel.SetDirectXConstantDirection(val)
		})

	directXItems := []fyne.CanvasObject{
		widget.NewLabel("DirectX FFB"),
		checkbDirectXInvertConstantForce,
	}

	if isPro {
		directXSpringStrengthSlider := components.Slider(
			uint16(effectSettings.DirectXSpringStrength),
			"Direct X spring forces strength (%)",
			100,
			func(f float64) {
				dev.Wheel.SetDirectXSpringStrength(byte(f))
			},
		)

		directXConstantStrengthSlider := components.Slider(
			uint16(effectSettings.DirectXConstantStrength),
			"Direct X constant forces strength (%)",
			100,
			func(f float64) {
				dev.Wheel.SetDirectXConstantStrength(byte(f))
			},
		)

		directXPeriodicStrengthSlider := components.Slider(
			uint16(effectSettings.DirectXPeriodicStrength),
			"Direct X periodic forces strength (%)",
			100,
			func(f float64) {
				dev.Wheel.SetDirectXPeriodicStrength(byte(f))
			},
		)

		directXItems = append(
			directXItems,
			directXSpringStrengthSlider,
			directXConstantStrengthSlider,
			directXPeriodicStrengthSlider,
		)
	}

	directXSettings := container.NewVBox(directXItems...)

	effectsContainer := container.NewVBox(
		commonSettings,
		directXSettings,
	)

	return effectsContainer
}
