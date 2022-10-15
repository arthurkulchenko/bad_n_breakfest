package models

import(
	"time"
)

type Reservation struct {
	FirstName string
	LastName string
	Email string
	Phone string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users struct {
	Id int
	FirstName string
	LastName string
	Email string
	Password string
	AccessLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Rooms struct {
	Id int
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restrictions struct {
	Id int
	RestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservations struct {
	Id int
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomId int
	Room Rooms
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoomRestrictions struct {
	Id int
	StartDate time.Time
	EndDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	RoomId int
	Room Rooms
	ReservationId int
	Reservation Reservations
	RestrictionId int
	Restriction Restrictions
}

