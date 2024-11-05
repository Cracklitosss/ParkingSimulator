package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FyneAdapter struct {
	window fyne.Window
	status *widget.Label
	logs   *widget.Label
}

func NewFyneAdapter() *FyneAdapter {
	a := app.New()
	w := a.NewWindow("Simulador de Estacionamiento")

	status := widget.NewLabel("Estado del estacionamiento")
	logs := widget.NewLabel("Log de Eventos")

	content := container.NewVBox(
		status,
		logs,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))

	return &FyneAdapter{
		window: w,
		status: status,
		logs:   logs,
	}
}

func (f *FyneAdapter) Show() {
	f.window.ShowAndRun()
}

func (f *FyneAdapter) UpdateStatus(text string) {
	f.status.SetText(text)
}

func (f *FyneAdapter) AddLog(text string) {
	currentText := f.logs.Text
	if len(currentText) > 1000 { // tamaño del log
		currentText = currentText[len(currentText)-1000:]
	}
	f.logs.SetText(currentText + "\n" + text)
}
