package RideBookingSystem

import "fmt"

type CustomerStatus interface {
	GetOnline(customer *Customer)
	GetOffline(customer *Customer)
	RequestRide(customer *Customer)
}

type OnlineStatus struct {
	status string
}

func (os *OnlineStatus) GetOnline(customer *Customer) {
	return
}

func (os *OnlineStatus) GetOffline(customer *Customer) {
	customer.customerStatus = &OfflineStatus{status: "Offline"}
}

func (os *OnlineStatus) RequestRide(customer *Customer) {
	fmt.Println("Ask Ride transition not valid")
}

type OfflineStatus struct {
	status string
}

func (os *OfflineStatus) GetOnline(customer *Customer) {
	fmt.Println("Get Online transition not valid")
}

func (os *OfflineStatus) GetOffline(customer *Customer) {
	return
}

func (os *OfflineStatus) RequestRide(customer *Customer) {
	customer.customerStatus = &RequestRideStatus{status: "RequestRide"}
}

type RequestRideStatus struct {
	status string
}

func (os *RequestRideStatus) GetOnline(customer *Customer) {
	fmt.Println("Get Online transition not valid")
	return
}

func (os *RequestRideStatus) GetOffline(customer *Customer) {
	customer.customerStatus = &OfflineStatus{status: "Offline"}
}

func (os *RequestRideStatus) RequestRide(customer *Customer) {
	return
}

type Customer struct {
	name           string
	id             string
	age            int
	isVerified     bool
	customerStatus CustomerStatus
}

type RideOrchestrator struct {
	DriverLocator DriverLocator
	DriverData    DriverData
	RidesData     RideData
}

func (cc *RideOrchestrator) RequestRide(pickUp Location, dropOff Location, customer *Customer) {
	customer.customerStatus.RequestRide(customer)
	driver := cc.DriverLocator.SelectDriver(pickUp, cc.DriverData.GetAvailableDrivers())
	if driver == nil {
		customer.customerStatus.GetOffline(customer)
		return
	}
	driver.status.InRide(driver)
	customer.customerStatus.GetOnline(customer)
	rd := &Ride{}
	rd.createRide(driver, customer, dropOff, pickUp)
	cc.RidesData.AddRiderData(*rd)
}

// Ensure triggered by driver
func (cc *RideOrchestrator) EndRide(ride Ride, currentLocation Location) {
	if ride.DropLocation.Latitude != currentLocation.Latitude || ride.DropLocation.Longitude != currentLocation.Longitude {
		return
	}
	cust := ride.Customer
	dri := ride.Driver
	cust.customerStatus.GetOffline(cust)
	dri.status.GetOnline(dri)
}
