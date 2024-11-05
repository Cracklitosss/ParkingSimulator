package ui

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Components struct {
	StartButton *widget.Button
	StopButton  *widget.Button
	Spaces      []*canvas.Text
}

func NewComponents(onStart func(), onStop func()) *Components {
	spaces := make([]*canvas.Text, 20)
	for i := range spaces {
		spaces[i] = canvas.NewText("Libre", color.Gray{})
	}

	return &Components{
		StartButton: widget.NewButton("Iniciar Simulación", onStart),
		StopButton:  widget.NewButton("Detener Simulación", onStop),
		Spaces:      spaces,
	}
}
