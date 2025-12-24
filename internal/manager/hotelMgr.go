package manager

import (
	"fmt"
	"hotel-mgmt-lld/internal/domain"
	"sync"
)

type HotelManager struct {
	CustomerMgr    *CustomerMgr
	PaymentMgr     *PaymentMgr
	ReservationMgr *ReservationMgr
}

var (
	HotelInstance *HotelManager
	HotelOnce     sync.Once
)

func NewHotelManager() *HotelManager {
	HotelOnce.Do(func() {
		HotelInstance = &HotelManager{
			CustomerMgr:    NewCustomerMgr(),
			PaymentMgr:     NewPaymentMgr(),
			ReservationMgr: NewReservationManager(),
		}
	})
	return HotelInstance
}

func (hm *HotelManager) BookRoom(paymentMethod domain.PaymentStrategy, customerId, roomId int) {
	paymentStrategy := hm.PaymentMgr.GetPaymentStrategy(paymentMethod)
	amount, err := hm.ReservationMgr.BookRoom(roomId, customerId)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	hm.PaymentMgr.Pay(paymentStrategy, amount)
}

func (hm *HotelManager) FindAvailableRooms(category ...domain.RoomCategory) {
	availableRooms := hm.ReservationMgr.FindAvailableRooms(category...)
	for _, v := range availableRooms {
		fmt.Printf("Room with id : %v available\n", v)
	}

}
