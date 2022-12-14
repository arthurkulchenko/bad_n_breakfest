package driver

import(
	"database/sql"
	"time"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := newDatabase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = d
	// err = testDB(d)
	// if err != nil {
	// 	return nil, error
	// }
	return dbConn, nil
}

func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil { return nil, err }
	if err = db.Ping(); err != nil { return nil, err }

	return db, nil
}

// func testDB(d * sql.DB) error {
// 	err := d.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

