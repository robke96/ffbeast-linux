package wheel

import (
	"encoding/binary"
	"errors"
	"time"
)

const stm32Poly uint32 = 0x04C11DB7

func crc32STM32(words ...uint32) uint32 {
	crc := uint32(0xFFFFFFFF)
	for _, w := range words {
		crc ^= w
		for i := 0; i < 32; i++ {
			if (crc & 0x80000000) != 0 {
				crc = (crc << 1) ^ stm32Poly
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}

func generateSerialKeys(deviceID [3]uint32) ([3]uint32, [3]uint32) {
	uid0 := deviceID[0]
	uid1 := deviceID[1]
	uid2 := deviceID[2]

	// Firmware path A: C = CRC(uid0,uid1,uid2)
	baseA := crc32STM32(uid0, uid1, uid2)
	keyA := [3]uint32{
		baseA ^ crc32STM32(uid0),
		baseA ^ crc32STM32(uid1),
		baseA ^ crc32STM32(uid2),
	}

	// Firmware path B: C = CRC(0x58F9)
	baseB := crc32STM32(0x000058F9)
	keyB := [3]uint32{
		baseB ^ crc32STM32(uid0),
		baseB ^ crc32STM32(uid1),
		baseB ^ crc32STM32(uid2),
	}

	return keyA, keyB
}

func (w *Wheel) sendActivationData(serialKey [3]uint32) error {
	if w.dev == nil {
		return errors.New("device not connected")
	}

	report := make([]byte, REPORT_LEN)
	report[0] = REPORT_GENERIC_INPUT_OUTPUT
	report[1] = DATA_FIRMWARE_ACTIVATION_DATA

	binary.LittleEndian.PutUint32(report[2:], serialKey[0])
	binary.LittleEndian.PutUint32(report[6:], serialKey[1])
	binary.LittleEndian.PutUint32(report[10:], serialKey[2])

	_, err := w.dev.Write(report)
	return err
}

func (w *Wheel) waitRegistrationState(registered bool) bool {
	for range 5 {
		time.Sleep(150 * time.Millisecond)
		licence := w.ReadFirmwareLicence()
		if licence == nil {
			continue
		}

		currentState := licence.IsRegistered != 0
		if currentState == registered {
			return true
		}
	}

	return false
}

func (w *Wheel) ActivateLicence() error {
	licence := w.ReadFirmwareLicence()
	if licence == nil {
		return errors.New("failed to read firmware licence")
	}

	keyA, keyB := generateSerialKeys(licence.DeviceId)

	if err := w.sendActivationData(keyA); err != nil {
		return err
	}
	if w.waitRegistrationState(true) {
		return nil
	}

	if err := w.sendActivationData(keyB); err != nil {
		return err
	}
	if w.waitRegistrationState(true) {
		return nil
	}

	return errors.New("activation failed: device stayed unregistered")
}

func (w *Wheel) DeactivateLicence() error {
	if err := w.sendActivationData([3]uint32{}); err != nil {
		return err
	}

	if w.waitRegistrationState(false) {
		return nil
	}

	return errors.New("deactivation failed: device stayed registered")
}
