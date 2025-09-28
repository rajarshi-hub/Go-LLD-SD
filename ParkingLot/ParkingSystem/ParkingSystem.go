package ParkingSystem

import (
	"fmt"
	"github.com/rajarshi-hub/ParkingLot/ParkingFloors"
	"github.com/rajarshi-hub/ParkingLot/ParkingPricing"
	"github.com/rajarshi-hub/ParkingLot/ParkingSlot"
	"github.com/rajarshi-hub/ParkingLot/Vehicle"
	"github.com/rajarshi-hub/ParkingLot/VehicleTicket"
	"github.com/rajarshi-hub/ParkingLot/VehicleType"
	"time"
)

type ParkingSystem struct {
	Floors         []*ParkingFloors.ParkingFloor
	TypeToPriceMap map[VehicleType.Type]ParkingPricing.ParkingPricing
}

func (ps *ParkingSystem) GetAvailableSlot(vehicleType VehicleType.Type) *ParkingSlot.ParkingSlot {
	for _, floor := range ps.Floors {
		slot := floor.GetAvailableSlot(vehicleType)
		if slot != nil {
			return slot
		}
	}
	return nil
}

func (ps *ParkingSystem) GenerateTicket(vehicle *Vehicle.Vehicle) *VehicleTicket.VehicleTicket {
	ts := time.Now()
	vehicleTicket := VehicleTicket.VehicleTicket{
		TicketId:      int(time.Now().UnixNano()),
		Vehicle:       vehicle,
		InTime:        &ts,
		OutTime:       nil,
		IsActive:      true,
		PriceStrategy: ps.TypeToPriceMap[vehicle.VehicleType],
	}
	return &vehicleTicket
}

func (ps *ParkingSystem) ParkVehicle(vehicle *Vehicle.Vehicle) *VehicleTicket.VehicleTicket {
	slot := ps.GetAvailableSlot(vehicle.VehicleType)
	if slot == nil {
		// Graceful error handling
		fmt.Printf("No slot available to park vehicle %v \n", vehicle.VehicleNo)
		return nil
	}
	vehicle.OccupiedSlot = slot
	fmt.Printf("Vehicle %v parked on %v \n", vehicle.VehicleNo, vehicle.OccupiedSlot.SId)
	return ps.GenerateTicket(vehicle)

}

func (ps *ParkingSystem) UnParkVehicle(vehicleTicket *VehicleTicket.VehicleTicket) {
	vehicle := vehicleTicket.Vehicle
	slot := vehicle.OccupiedSlot
	vehicle.Exit()
	slot.FreeAvailableSlot()
	vehicleTicket.MarkResolved()
}
