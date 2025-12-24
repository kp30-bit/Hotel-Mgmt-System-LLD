package domain

type Room struct {
	Id         int
	IsBooked   bool
	Category   RoomCategory // scale 1-3 where 3 is best quality room
	CustomerId int
}
