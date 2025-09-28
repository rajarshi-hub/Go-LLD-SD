package VehicleType

type Type int

//type SlotType struct {
//	VehicleType VehicleType.Type
//	HasCharger  bool
//}

const (
	Truck Type = iota
	SUV
	Hatchback
	Bike
	Scooter
)
