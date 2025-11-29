package RideBookingSystem

import "math"

type CustomerData struct {
	customers []*Customer
}

type DriverData struct {
	drivers []*Driver
}

func (dd *DriverData) GetAvailableDrivers() []*Driver {
	// Logic to exclude unavailable ones
	var availableDrivers []*Driver
	for _, driver := range dd.drivers {
		if driver.status.GetStatus() == "Online" {
			availableDrivers = append(availableDrivers, driver)
		}
	}
	return availableDrivers
}

type RideData struct {
	Rides []Ride
}

func (rd *RideData) AddRiderData(ride Ride) {
	rd.Rides = append(rd.Rides, ride)
}

// TODO: Add, remove, update functions here for all data layers

type DriverLocator struct{}

func (db *DriverLocator) SelectDriver(pickUp Location, drivers []*Driver) *Driver {
	// Optimal driver
	minDist := math.MaxFloat64
	var proposedDriver *Driver
	for _, driver := range drivers {
		lat := driver.location.Latitude
		long := driver.location.Longitude
		dist := math.Sqrt((pickUp.Latitude-lat)*(pickUp.Latitude-lat) + (pickUp.Longitude-long)*(pickUp.Longitude-long))
		if dist < minDist {
			proposedDriver = driver
			minDist = dist
		}
	}
	return proposedDriver
}
