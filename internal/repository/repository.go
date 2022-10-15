package repository

import(
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) (int, error)
}
