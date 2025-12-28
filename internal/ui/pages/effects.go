package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/robke96/ffbeast-linux/internal/device"
	"github.com/robke96/ffbeast-linux/internal/ui/components"
)

func EffectsPage(dev *device.Device) *fyne.Container {
	effectSettings := dev.Wheel.ReadEffectSettings()

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

	commonSettings := container.NewVBox(
		widget.NewLabel("Common"),
		motionRangeSlider,
		totalStrengthSlider,
		staticDampeningSlider,
	)

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

	directXSettings := container.NewVBox(
		widget.NewLabel("DirectX FFB"),
		checkbDirectXInvertConstantForce,
	)

	effectsContainer := container.NewVBox(
		commonSettings,
		directXSettings,
	)

	return effectsContainer
}
