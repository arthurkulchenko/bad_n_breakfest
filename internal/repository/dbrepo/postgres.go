package dbrepo

import(
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"time"
	"context"
)

func (m * postgresDBRepo) AllUsers() bool {
	return true
}

func (m * postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	statement := `
	  insert into reservations
			(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id
	`
	var recordId int
	err := m.DB.QueryRowContext(
		ctx,
		statement,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	).Scan(&recordId)
	if err != nil {
		return 0, err
	}
	return int(recordId), nil
}

func (m * postgresDBRepo) InsertRoomRestriction(res models.RoomRestriction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	statement := `
	  insert into room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
		values ($1,$2,$3,$4,$5,$6,$7) returning id
	`
	var recordId int
	err := m.DB.QueryRowContext(
		ctx,
		statement,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		res.ReservationId,
		res.RestrictionId,
		time.Now(),
		time.Now(),
	).Scan(&recordId)
	if err != nil {
		return 0, err
	}
	return int(recordId), nil
}
