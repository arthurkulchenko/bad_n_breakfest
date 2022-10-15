package dbrepo

import(
	"database/sql"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{ App: a, DB: conn, }
}
