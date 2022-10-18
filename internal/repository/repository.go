package repository

import(
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"time"
)

type DatabaseInterface interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) (int, error)
	SearchAvailabilityByDatedCount(start, end *time.Time, roomId int) (int, error)
	SerachRoomsAvailability(start, end *time.Time) ([]models.Room, error)
	FindRoomById(id int) (models.Room, error)
}
