package main

import(
	"github.com/arthurkulchenko/bed_n_breakfest/pkg/config"
	"github.com/arthurkulchenko/bed_n_breakfest/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routes(appP *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.RepositoryPointer.Home)
	mux.Get("/about", handlers.RepositoryPointer.About)
	mux.Get("/generals", handlers.RepositoryPointer.General)
	mux.Get("/majors", handlers.RepositoryPointer.Major)
	mux.Get("/reservation", handlers.RepositoryPointer.Reservation)
	mux.Get("/contacts", handlers.RepositoryPointer.Contact)
	mux.Get("/search-availability", handlers.RepositoryPointer.SearchAvailability)
	mux.Post("/search-availability", handlers.RepositoryPointer.PostSearchAvailability)
	mux.Get("/search-availability-json", handlers.RepositoryPointer.GetSearchAvailabilityJson)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
