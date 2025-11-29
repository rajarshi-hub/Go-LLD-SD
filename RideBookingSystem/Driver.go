package RideBookingSystem

type GenderType string

type DriverStatus interface {
	GetOnline(driver *Driver)
	GetOffline(driver *Driver)
	InRide(driver *Driver)
	GetStatus() string
}

type DOnlineStatus struct {
	status string
}

func (os *DOnlineStatus) GetOnline(driver *Driver) {
	return
}

func (os *DOnlineStatus) GetOffline(driver *Driver) {
	driver.status = &DOfflineStatus{status: "Offline"}
}

func (os *DOnlineStatus) InRide(driver *Driver) {
	driver.status = &DInRideStatus{status: "InRide"}
}

func (os *DOnlineStatus) GetStatus() string {
	return os.status
}

type DOfflineStatus struct {
	status string
}

func (os *DOfflineStatus) GetOnline(driver *Driver) {
	driver.status = &DOnlineStatus{status: "Online"}
}

func (os *DOfflineStatus) GetOffline(driver *Driver) {
	return
}

func (os *DOfflineStatus) InRide(driver *Driver) {
	return
}

func (os *DOfflineStatus) GetStatus() string {
	return os.status
}

type DInRideStatus struct {
	status string
}

func (os *DInRideStatus) GetOnline(driver *Driver) {
	driver.status = &DOnlineStatus{status: "Online"}
}

func (os *DInRideStatus) GetOffline(driver *Driver) {
	return
}

func (os *DInRideStatus) InRide(driver *Driver) {
	return
}

func (os *DInRideStatus) GetStatus() string {
	return os.status
}

type Driver struct {
	name      string
	vehicleNo string
	age       int
	gender    GenderType
	status    DriverStatus
	location  Location
}

func (os *Driver) GetOnline() {
	os.status.GetOnline(os)
}

func (os *Driver) GetOffline() {
	os.status.GetOffline(os)
}
