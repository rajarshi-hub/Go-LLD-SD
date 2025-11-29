package RideBookingSystem

type Ride struct {
	Driver         *Driver
	Customer       *Customer
	DropLocation   Location
	PickUpLocation Location
	Fare           int
}

func (rd *Ride) createRide(driver *Driver, Customer *Customer, dropLocation Location, pickUpLocation Location) {
	rd.Driver = driver
	rd.Customer = Customer
	rd.DropLocation = dropLocation
	rd.PickUpLocation = pickUpLocation
}
