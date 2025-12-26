package ui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func EffectsPage() *widget.Card {
	button := widget.NewButton("Click", func() {
		fmt.Println("hi")
	})

	card := widget.NewCard("Common", "", button)
	return card
}
