package wheel

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (w *Wheel) ReadData(reportFeature byte, dest interface{}) error {
	if w.dev == nil {
		fmt.Println("device not connected")
		return nil
	}

	report := make([]byte, 65)
	report[0] = reportFeature

	result, err := w.dev.GetFeatureReport(report)
	if err != nil || result <= 0 {
		fmt.Println("problem")
		return nil
	}

	err = binary.Read(
		bytes.NewReader(report[1:]),
		binary.LittleEndian,
		dest,
	)
	if err != nil {
		return fmt.Errorf("binary read failed: %w", err)
	}

	return nil
}

func (w *Wheel) ReadEffectSettings() *EffectSettings {
	effect := &EffectSettings{}
	err := w.ReadData(REPORT_EFFECT_SETTINGS_FEATURE, effect)
	if err != nil {
		fmt.Println("Error reading effect settings:", err)
		return nil
	}
	return effect
}

func (w *Wheel) ReadHardwareSettings() *HardwareSettings {
	hardware := &HardwareSettings{}
	err := w.ReadData(REPORT_HARDWARE_SETTINGS_FEATURE, hardware)
	if err != nil {
		fmt.Println("Error reading hardware settings:", err)
		return nil
	}
	return hardware
}
