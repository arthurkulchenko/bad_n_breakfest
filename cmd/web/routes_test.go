package main

import(
	"testing"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
)

// Test return type
func TestRoutes(test *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		test.Error(fmt.Sprintf("type is not *chi.Mux, it is %T", v))
	}
}
