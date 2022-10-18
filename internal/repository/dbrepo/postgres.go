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

func (m * postgresDBRepo) SearchAvailabilityByDatedCount(start, end *time.Time, roomId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	// params = ['2022-10-15', '2022-10-30']
	selectStatement := `
		select count(id) from room_restrictions where start_date < $1 and $2 < end_date and room_id = $3
	`
	var count int
	err := m.DB.QueryRowContext(
		ctx,
		selectStatement,
		start,
		end,
		roomId,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (m *postgresDBRepo) SerachRoomsAvailability(start, end *time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	query := `
		select id, room_name from rooms where id not in (select room_id from room_restrictions where start_date < $1 and $2 < end_date)
	`
	var rooms []models.Room
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil { return rooms, err }
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.Id, &room.RoomName,)
		rooms = append(rooms, room)
		if err != nil { return rooms, err }
	}
	if err = rows.Err(); err != nil { return rooms, err }
	return rooms, nil
}

func (m *postgresDBRepo) FindRoomById(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	var room models.Room
	query := `select * from rooms where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.Id,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return room, err
	}
	return room, nil
}

// func (m *postgresDBRepo) FindById(entity string, id int) (interface{}, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
// 	defer cancel()
// 	var modelInstance interface{}
// 	query := `select * from $1 where id = $2`
// 	row, err := m.DB.QueryRowContext(ctx, query, string, id)
// 	err := row.Scan()
// }
