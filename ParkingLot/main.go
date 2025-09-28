package main

import (
	"fmt"
	"github.com/rajarshi-hub/ParkingLot/ParkingFloors"
	"github.com/rajarshi-hub/ParkingLot/ParkingPricing"
	"github.com/rajarshi-hub/ParkingLot/ParkingSlot"
	"github.com/rajarshi-hub/ParkingLot/ParkingSystem"
	"github.com/rajarshi-hub/ParkingLot/Vehicle"
	"github.com/rajarshi-hub/ParkingLot/VehicleType"
	"time"
)

func main() {
	slot1 := ParkingSlot.ParkingSlot{
		SId:             0,
		IsAvailable:     true,
		VehicleTypeSlot: 1,
	}
	slot2 := ParkingSlot.ParkingSlot{
		SId:             1,
		IsAvailable:     true,
		VehicleTypeSlot: 2,
	}
	slot3 := ParkingSlot.ParkingSlot{
		SId:             2,
		IsAvailable:     true,
		VehicleTypeSlot: 2,
	}
	slot4 := ParkingSlot.ParkingSlot{
		SId:             6,
		IsAvailable:     true,
		VehicleTypeSlot: 3,
	}
	floor1 := ParkingFloors.ParkingFloor{Slots: []*ParkingSlot.ParkingSlot{&slot1, &slot2}}
	floor2 := ParkingFloors.ParkingFloor{Slots: []*ParkingSlot.ParkingSlot{&slot3}}
	floor3 := ParkingFloors.ParkingFloor{Slots: []*ParkingSlot.ParkingSlot{&slot4}}
	parkingSystem := ParkingSystem.ParkingSystem{Floors: []*ParkingFloors.ParkingFloor{&floor1, &floor2, &floor3}}
	parkingSystem.TypeToPriceMap = map[VehicleType.Type]ParkingPricing.ParkingPricing{
		1: &ParkingPricing.CarParkingPricing{
			BasePrice:    50,
			PricePerHour: 20,
		},
		0: &ParkingPricing.CarParkingPricing{
			BasePrice:    50,
			PricePerHour: 20,
		},
		2: &ParkingPricing.CarParkingPricing{
			BasePrice:    50,
			PricePerHour: 20,
		},
		3: &ParkingPricing.BikeParkingPricing{
			BasePrice:    20,
			PricePerHour: 10,
		},
		4: &ParkingPricing.BikeParkingPricing{
			BasePrice:    20,
			PricePerHour: 10,
		},
	}
	vehicle1 := Vehicle.Vehicle{
		VehicleType:  1,
		VehicleNo:    "UP78-1234",
		OccupiedSlot: nil,
	}
	vehicle2 := Vehicle.Vehicle{
		VehicleType:  2,
		VehicleNo:    "UP78-1256",
		OccupiedSlot: nil,
	}
	vehicle3 := Vehicle.Vehicle{
		VehicleType:  2,
		VehicleNo:    "UP78-1289",
		OccupiedSlot: nil,
	}
	vehicle4 := Vehicle.Vehicle{
		VehicleType:  2,
		VehicleNo:    "UP78-1345",
		OccupiedSlot: nil,
	}
	vehicle5 := Vehicle.Vehicle{
		VehicleType:  3,
		VehicleNo:    "UP78-1678",
		OccupiedSlot: nil,
	}
	ticket1 := parkingSystem.ParkVehicle(&vehicle1)
	ticket2 := parkingSystem.ParkVehicle(&vehicle2)
	ticket3 := parkingSystem.ParkVehicle(&vehicle3)
	ticket4 := parkingSystem.ParkVehicle(&vehicle4)
	ticket5 := parkingSystem.ParkVehicle(&vehicle5)

	time.Sleep(5 * time.Second)

	parkingSystem.UnParkVehicle(ticket1)
	parkingSystem.UnParkVehicle(ticket2)
	parkingSystem.UnParkVehicle(ticket5)
	parkingSystem.UnParkVehicle(ticket3)

	ticket4 = parkingSystem.ParkVehicle(&vehicle4)

	fmt.Println(&ticket1, &ticket2, &ticket3, &ticket4)

}
