package domain

type RoomCategory int

const (
	Suite RoomCategory = iota + 1
	Deluxe
	Premiere
)

type PaymentStrategy int

const (
	CardPayment PaymentStrategy = iota
	UpiPayment
)
