package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/arthurkulchenko/bed_n_breakfest/pkg/config"
	"github.com/arthurkulchenko/bed_n_breakfest/pkg/handlers"
	"log"
	"net/http"
	"time"
)

const PORT_NUMBER = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.Env = "development"
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Env == "production"
	app.Session = session

	templateCache, err := handlers.CreateTemplateCache()
	if err != nil { log.Fatal("Cannot create template cache") }

	app.TemplateCache = templateCache
	app.PortNumber = PORT_NUMBER
	app.UseCache = false

	handlers.SetConfigAndRepository(&app)

	// handlersRepo := handlers.NewRepo(&app)
	// handlers.NewHandlers(handlersRepo)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", app.PortNumber))
	server := &http.Server { Addr: app.PortNumber, Handler: Routes(&app) }
	err = server.ListenAndServe()
	log.Fatal(err)
}
