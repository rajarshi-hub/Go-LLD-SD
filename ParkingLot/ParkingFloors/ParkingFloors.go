package ParkingFloors

import (
	"github.com/rajarshi-hub/ParkingLot/ParkingSlot"
	"github.com/rajarshi-hub/ParkingLot/VehicleType"
)

type ParkingFloor struct {
	fid   int
	Slots []*ParkingSlot.ParkingSlot
}

func (pf *ParkingFloor) GetAvailableSlot(vehicleType VehicleType.Type) *ParkingSlot.ParkingSlot {
	for _, slot := range pf.Slots {
		if slot.IsAvailable && slot.VehicleTypeSlot == vehicleType {
			slot.IsAvailable = false
			return slot
		}
	}
	return nil

}
