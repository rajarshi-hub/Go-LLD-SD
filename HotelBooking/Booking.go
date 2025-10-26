package HotelBooking

import (
	"fmt"
	"time"
)

type BookingStatus int

const Booked BookingStatus = 0
const Canceled BookingStatus = 1

type DateRange struct {
	startDate time.Time
	endDate   time.Time
}

type User struct {
	username   string
	userId     int
	isVerified bool
}

type Booking struct {
	Id            int
	User          User
	bookingStatus BookingStatus
	bookedRooms   map[int]DateRange
}

func (bm *BookingManager) CancelBooking(bookingId int) error {
	for i := range bm.bookings {
		if bm.bookings[i].Id == bookingId {
			if bm.bookings[i].bookingStatus == Canceled {
				return fmt.Errorf("booking already canceled")
			}
			bm.bookings[i].bookingStatus = Canceled
			return nil
		}
	}
	return nil
}

type BookingManager struct {
	bookings []Booking
}

func (bm *BookingManager) CheckOverlapping(room Room, dateRange DateRange) bool {
	isAvailable := true
	for _, booking := range bm.bookings {
		if booking.bookingStatus == Canceled {
			continue
		}
		for roomIDCheck, dateCheck := range booking.bookedRooms {
			if room.roomID != roomIDCheck {
				continue
			}
			if dateRange.startDate.After(dateCheck.endDate) || dateCheck.startDate.After(dateRange.endDate) {
				continue
			}
			isAvailable = false
		}
	}
	return isAvailable
}

func (bm *BookingManager) QueryBooking(avail map[RoomType]int, dateRange DateRange, rm *RoomManager) (bool, error) {
	// TODO: Partial availability: if someone requests 3 rooms but only 2 are available
	for roomType, quant := range avail {
		availableRooms := 0
		allRooms := rm.GetRooms()
		for _, room := range allRooms {
			if room.roomType != roomType {
				continue
			}
			if bm.CheckOverlapping(room, dateRange) {
				availableRooms++
			}
		}
		if availableRooms < quant {
			return false, nil
		}
	}
	return true, nil
}

func (bm *BookingManager) MakeBooking(rooms []Room, dateRange DateRange, id int) (*Booking, error) {
	// TODO: handling concurrency
	// TODO: Later: maintain date-based availability maps.
	// TODO: Per-hour booking support — use time.Time precision instead of day-level.
	// TODO: Validations as QueryBooking before actually bokking
	booking := Booking{}
	booking.bookingStatus = Booked
	booking.Id = id
	booking.User = User{
		userId: id,
	}
	booking.bookedRooms = make(map[int]DateRange)
	for _, room := range rooms {
		booking.bookedRooms[room.roomID] = dateRange
	}
	bm.bookings = append(bm.bookings, booking)
	return &booking, nil
}
