package main

import (
	"hotel-mgmt-lld/internal/domain"
	"hotel-mgmt-lld/internal/manager"
)

func main() {
	room1 := &domain.Room{
		Id:       1,
		IsBooked: false,
		Category: domain.Suite,
	}
	room2 := &domain.Room{
		Id:       2,
		IsBooked: false,
		Category: domain.Suite,
	}
	room3 := &domain.Room{
		Id:       3,
		IsBooked: false,
		Category: domain.Premiere,
	}
	hotelMgr := manager.NewHotelManager()
	hotelMgr.ReservationMgr.AddRoom(room1)
	hotelMgr.ReservationMgr.AddRoom(room2)
	hotelMgr.ReservationMgr.AddRoom(room3)
	cust1 := domain.Customer{
		UserId: 1,
	}
	hotelMgr.CustomerMgr.AddCustomer(&cust1)
	hotelMgr.BookRoom(domain.CardPayment, 1, 1)
	hotelMgr.FindAvailableRooms(domain.Suite)
}
