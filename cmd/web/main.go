package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/handlers"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/helpers"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/driver"
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
	dbConn, err := run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	defer dbConn.SQL.Close()
	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", app.PortNumber))
	server := &http.Server { Addr: app.PortNumber, Handler: Routes(&app) }
	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation {})
	gob.Register(models.User {})
	gob.Register(models.Room {})
	gob.Register(models.Restriction {})
	app.Env = "development"
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile )
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Env == "production"
	app.Session = session
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=breakfest_and_bed user=wayfarer password=wayfarer")
	if err != nil {
		log.Fatal("Cannot connect db.")
		return nil, err
	}

	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.PortNumber = PORT_NUMBER
	app.UseCache = false

	handlers.SetConfigAndRepository(&app, db)
	helpers.NewHelpers(&app)
	return db, nil
}
