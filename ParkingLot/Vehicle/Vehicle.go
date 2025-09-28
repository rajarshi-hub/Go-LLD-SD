package Vehicle

import (
	"fmt"
	"github.com/rajarshi-hub/ParkingLot/ParkingSlot"
	"github.com/rajarshi-hub/ParkingLot/VehicleType"
)

type Vehicle struct {
	VehicleType  VehicleType.Type
	VehicleNo    string
	OccupiedSlot *ParkingSlot.ParkingSlot
}

func (v *Vehicle) Exit() {
	v.OccupiedSlot = nil
	fmt.Printf("Vehicle Exit for %v \n", v.VehicleNo)
}
