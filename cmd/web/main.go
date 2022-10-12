package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/handlers"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/helpers"
	"log"
	"net/http"
	"time"
	"encoding/gob"
	"os"
)

const PORT_NUMBER = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	gob.Register(models.Reservation {})
	app.Env = "development"
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile )
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Env == "production"
	app.Session = session

	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.PortNumber = PORT_NUMBER
	app.UseCache = false

	handlers.SetConfigAndRepository(&app)
	helpers.NewHelpers(&app)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", app.PortNumber))
	server := &http.Server { Addr: app.PortNumber, Handler: Routes(&app) }
	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	return nil
}
