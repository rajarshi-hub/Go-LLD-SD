package ParkingSlot

import (
	"github.com/rajarshi-hub/ParkingLot/VehicleType"
)

type ParkingSlot struct {
	SId             int
	IsAvailable     bool
	VehicleTypeSlot VehicleType.Type
}

func (ps *ParkingSlot) FreeAvailableSlot() {
	ps.IsAvailable = true
}
