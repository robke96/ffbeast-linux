package ui

import (
	"image"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/disintegration/imaging"
	_ "golang.org/x/image/webp"
)

type WheelIcon struct {
	object      fyne.CanvasObject
	imageWidget *canvas.Image
	base        image.Image
	motionRange float64
}

func NewWheelIcon(path string, size float32) *WheelIcon {
	base := decodeWheelImage(path)
	imageWidget := canvas.NewImageFromImage(base)
	imageWidget.FillMode = canvas.ImageFillContain

	wrapped := container.NewGridWrap(
		fyne.NewSize(size, size),
		imageWidget,
	)

	return &WheelIcon{
		object:      wrapped,
		imageWidget: imageWidget,
		base:        base,
		motionRange: 900,
	}
}

func (w *WheelIcon) Object() fyne.CanvasObject {
	return w.object
}

func (w *WheelIcon) SetPosition(position int16) {
	// Position is normalized to [-10000, 10000].
	angle := -((float64(position) / 10000.0) * (w.motionRange / 2.0))

	rotated := imaging.Rotate(w.base, angle, color.Transparent)
	bw := w.base.Bounds().Dx()
	bh := w.base.Bounds().Dy()
	if rotated.Bounds().Dx() != bw || rotated.Bounds().Dy() != bh {
		rotated = imaging.CropCenter(rotated, bw, bh)
	}

	fyne.Do(func() {
		w.imageWidget.Image = rotated
		w.imageWidget.Refresh()
	})
}

func (w *WheelIcon) SetMotionRange(motionRange uint16) {
	if motionRange == 0 {
		w.motionRange = 900
		return
	}
	w.motionRange = float64(motionRange)
}

func decodeWheelImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		return image.NewNRGBA(image.Rect(0, 0, 1, 1))
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return image.NewNRGBA(image.Rect(0, 0, 1, 1))
	}

	return img
}
