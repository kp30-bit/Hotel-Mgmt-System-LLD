package manager

import (
	"fmt"
	"hotel-mgmt-lld/internal/domain"
	"sync"
)

type ReservationMgr struct {
	RoomList map[int]*domain.Room
}

var (
	ReservationMgrInstance *ReservationMgr
	ReservationMgrOnce     sync.Once
)

func NewReservationManager() *ReservationMgr {
	ReservationMgrOnce.Do(func() {
		ReservationMgrInstance = &ReservationMgr{
			RoomList: make(map[int]*domain.Room),
		}
	})
	return ReservationMgrInstance
}

func (rm *ReservationMgr) FindAvailableRooms(category ...domain.RoomCategory) []int { //category can be treated as slice of ints and we can call f() , f(1) or f(1,2)
	var resultList []int

	for _, room := range rm.RoomList {
		if room.IsBooked {
			continue
		}

		if len(category) == 0 {
			resultList = append(resultList, room.Id)
		} else if len(category) == 1 && category[0] == room.Category {
			resultList = append(resultList, room.Id)
		}
	}
	return resultList
}

func (rm *ReservationMgr) BookRoom(roomId, customerId int) (int, error) {
	room, exists := rm.RoomList[roomId]
	if !exists {
		return -1, fmt.Errorf("Room with room id : %v is not available\n", room.Id)
	}
	room.IsBooked = true
	room.CustomerId = customerId
	return int(room.Category) * 1000, nil
}

func (rm *ReservationMgr) CancelRoomBooking(roomId, customerId int) error {
	room, exists := rm.RoomList[roomId]
	if !exists {
		return fmt.Errorf("Room with room id: %v does not exist\n", room.Id)
	}
	if room.IsBooked && room.CustomerId == customerId {
		room.IsBooked = false
		room.CustomerId = -1
	}
	return nil
}

func (rm *ReservationMgr) AddRoom(room *domain.Room) {
	rm.RoomList[room.Id] = room
}
