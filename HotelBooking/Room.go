package HotelBooking

type Room struct {
	roomID   int
	roomType RoomType
}

type RoomType struct {
	roomTypeID   int
	roomTypeName string
}

type RoomManager struct {
	rooms []Room
}

func (rm *RoomManager) AddRooms(room Room) {
	rm.rooms = append(rm.rooms, room)
}

func (rm *RoomManager) GetRooms() []Room {
	return rm.rooms
}
