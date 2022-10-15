package models

import(
	"time"
)

// type Reservation struct {
// 	FirstName string
// 	LastName string
// 	Email string
// 	Phone string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type User struct {
	Id int
	FirstName string
	LastName string
	Email string
	Password string
	AccessLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Room struct {
	Id int
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restriction struct {
	Id int
	RestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservation struct {
	Id int
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomId int
	Room Room
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoomRestriction struct {
	Id int
	StartDate time.Time
	EndDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	RoomId int
	Room Room
	ReservationId int
	Reservation Reservation
	RestrictionId int
	Restriction Restriction
}

