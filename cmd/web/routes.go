package main

import(
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routes(appP *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.ControllerPointer.Home)
	mux.Get("/about", handlers.ControllerPointer.About)
	mux.Get("/generals", handlers.ControllerPointer.General)
	mux.Get("/majors", handlers.ControllerPointer.Major)
	mux.Get("/reservation", handlers.ControllerPointer.Reservation)
	mux.Post("/reservation", handlers.ControllerPointer.PostReservation)
	mux.Get("/contacts", handlers.ControllerPointer.Contact)
	mux.Get("/search-availability", handlers.ControllerPointer.SearchAvailability)
	mux.Post("/search-availability", handlers.ControllerPointer.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.ControllerPointer.PostSearchAvailabilityJson)
	mux.Get("/reservation-summary", handlers.ControllerPointer.GetReservationSummary)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
