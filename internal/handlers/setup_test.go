package handlers

import(
	"net/http"
	"time"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/alexedwards/scs/v2"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"log"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {
	gob.Register(models.Reservation {})
	app.Env = "test"
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Env == "production"
	app.Session = session

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.PortNumber = "PORT_NUMBER"
	app.UseCache = false

	SetConfigAndRepository(&app)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", RepositoryPointer.Home)
	mux.Get("/about", RepositoryPointer.About)
	mux.Get("/generals", RepositoryPointer.General)
	mux.Get("/majors", RepositoryPointer.Major)
	mux.Get("/reservation", RepositoryPointer.Reservation)
	mux.Post("/reservation", RepositoryPointer.PostReservation)
	mux.Get("/contacts", RepositoryPointer.Contact)
	mux.Get("/search-availability", RepositoryPointer.SearchAvailability)
	mux.Post("/search-availability", RepositoryPointer.PostSearchAvailability)
	mux.Post("/search-availability-json", RepositoryPointer.PostSearchAvailabilityJson)
	mux.Get("/reservation-summary", RepositoryPointer.GetReservationSummary)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
