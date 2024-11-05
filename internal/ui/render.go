package ui

import (
	"fmt"
	"parking-simulator/pkg/simulation"
	"time"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Renderer struct {
	Adapter    *FyneAdapter
	Simulator  *simulation.Simulator
	Components *Components
	stopChan   chan struct{}
}

func NewRenderer(adapter *FyneAdapter, simulator *simulation.Simulator) *Renderer {
	r := &Renderer{
		Adapter:    adapter,
		Simulator:  simulator,
		Components: NewComponents(nil, nil),
		stopChan:   make(chan struct{}),
	}

	r.Components.StartButton.OnTapped = r.StartSimulation
	r.Components.StopButton.OnTapped = r.StopSimulation

	canvasObjects := make([]fyne.CanvasObject, len(r.Components.Spaces))
	for i, space := range r.Components.Spaces {
		canvasObjects[i] = space
	}

	adapter.window.SetContent(
		container.NewVBox(
			adapter.status,
			container.NewGridWithColumns(5, canvasObjects...),
			r.Components.StartButton,
			r.Components.StopButton,
		),
	)

	return r
}

func (r *Renderer) StartSimulation() {
	go func() {
		for i := 0; i < r.Simulator.VehicleCount; i++ {
			select {
			case <-r.stopChan:
				return
			default:
				vehicle := simulation.NewVehicle(i)
				go r.handleVehicle(vehicle)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func (r *Renderer) StopSimulation() {
	r.stopChan <- struct{}{}
	r.Adapter.AddLog("Simulación detenida.")
}

func (r *Renderer) handleVehicle(vehicle *simulation.Vehicle) {
	for !r.Simulator.ParkingLot.EnterVehicle(vehicle) {
		r.Adapter.AddLog(fmt.Sprintf("Vehículo %d esperando por espacio...", vehicle.ID))
		time.Sleep(time.Second)
	}

	r.Components.Spaces[vehicle.ID%20].Text = "Ocupado"
	r.Components.Spaces[vehicle.ID%20].Color = color.RGBA{255, 0, 0, 255}
	r.Components.Spaces[vehicle.ID%20].Refresh()

	r.Adapter.UpdateStatus(fmt.Sprintf("Vehículo %d entró al estacionamiento", vehicle.ID))
	r.Adapter.AddLog(fmt.Sprintf("Vehículo %d estacionado.", vehicle.ID))

	vehicle.StayParked()

	r.Simulator.ParkingLot.ExitVehicle(vehicle)
	r.Components.Spaces[vehicle.ID%20].Text = "Libre"
	r.Components.Spaces[vehicle.ID%20].Color = color.Gray{}
	r.Components.Spaces[vehicle.ID%20].Refresh()

	r.Adapter.UpdateStatus(fmt.Sprintf("Vehículo %d salió del estacionamiento", vehicle.ID))
	r.Adapter.AddLog(fmt.Sprintf("Vehículo %d ha salido del estacionamiento.", vehicle.ID))
}
