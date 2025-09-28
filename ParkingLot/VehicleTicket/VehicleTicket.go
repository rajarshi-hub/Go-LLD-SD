package VehicleTicket

import (
	"fmt"
	"github.com/rajarshi-hub/ParkingLot/ParkingPricing"
	"github.com/rajarshi-hub/ParkingLot/Vehicle"
	"time"
)

type VehicleTicket struct {
	TicketId      int
	Vehicle       *Vehicle.Vehicle
	InTime        *time.Time
	OutTime       *time.Time
	IsActive      bool
	PriceStrategy ParkingPricing.ParkingPricing
}

func (vt *VehicleTicket) MarkResolved() {
	vt.IsActive = false
	ts := time.Now()
	vt.OutTime = &ts
	price := vt.PriceStrategy.CalculatePrice(vt.InTime, vt.OutTime)
	fmt.Printf("Vehicle MarkResolved for %v at %v and has to pay price %v \n", vt.Vehicle.VehicleNo, vt.OutTime, price)
}
