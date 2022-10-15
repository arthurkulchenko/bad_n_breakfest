package handlers

import (
	"fmt"
	// "errors"
	"encoding/json"
	"net/http"
	"html/template"
	"log"
	"path/filepath"
	"bytes"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"github.com/justinas/nosurf"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/models"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/forms"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/helpers"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/repository"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/repository/dbrepo"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/driver"
)

var appConfigP *config.AppConfig
var RepositoryPointer *Repository
// RepositoryPointer.AppConfigPointer => *config.AppConfig

type Repository struct {
	AppConfigPointer *config.AppConfig
	DB repository.DatabaseRepo
}

func SetConfigAndRepository(appConfigPointer *config.AppConfig, db *driver.DB) {
	appConfigP = appConfigPointer
	RepositoryPointer = &Repository { AppConfigPointer: appConfigPointer, DB: dbrepo.NewPostgresRepo(db.SQL, appConfigPointer), }
}

// func NewRepo(pointer *config.AppConfig, db *driver.DB) *Repository {
// 	return &Repository {
// 		AppConfigPointer: pointer,
// 		DB dbrepo.NewPostgresRepo(db.SQL, a)
// 	}
// }

// func NewHandlers(repositoryPointer *Repository) {
// 	RepositoryPointer = repositoryPointer
// }

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Watrning string
	Error string
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

func (receiver *Repository) Home(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	session := receiver.AppConfigPointer.Session

	var users string = fmt.Sprintf("%v", receiver.DB.AllUsers())
	stringMap["remoteaddr"] = users

	stringMap["remoteaddr1"] = request.RemoteAddr
	session.Put(request.Context(), "remoteaddr", request.RemoteAddr)

	renderTemplate(response, request, "home.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) About(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	session := receiver.AppConfigPointer.Session
	stringMap["remoteaddr"] = session.GetString(request.Context(), "remoteaddr")

	renderTemplate(response, request, "about.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Reservation(response http.ResponseWriter, request *http.Request) {
	var nullReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = nullReservation

	renderTemplate(response, request, "reservation.page.tmpl", &models.TemplateData { Form: forms.New(nil), Data: data })
}

func (receiver *Repository) PostReservation(response http.ResponseWriter, request *http.Request) {
	formError := request.ParseForm()
	if formError != nil {
		helpers.ServerError(response, formError)
		return
	}

	reservation := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName: request.Form.Get("last_name"),
		Email: request.Form.Get("email"),
		Phone: request.Form.Get("phone"),
	}
	form := forms.New(request.PostForm)
	// form.Has("first_name", request)
	form.Required("first_name", "last_name", "email")
	form.MinLen(3, "first_name", "last_name", "email")
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renderTemplate(response, request, "reservation.page.tmpl", &models.TemplateData { Form: form, Data: data })
		return
	}

	receiver.AppConfigPointer.Session.Put(request.Context(), "reservation", reservation)
	http.Redirect(response, request, "/reservation-summary", http.StatusSeeOther)
}

func (receiver *Repository) GetReservationSummary(response http.ResponseWriter, request *http.Request) {
	// stringMap := make(map[string]string)
	data := make(map[string]interface{})
	session := receiver.AppConfigPointer.Session
	reservation, fetchingStatus := session.Get(request.Context(), "reservation").(models.Reservation)
	if !fetchingStatus {
		receiver.AppConfigPointer.ErrorLog.Println("Can't get error from session")
		session.Put(request.Context(), "error", "Can't get reservation")
		http.Redirect(response, request, "/reservation", http.StatusTemporaryRedirect)
		session.Remove(request.Context(), "reservation")
		return
	}
	data["reservation"] = reservation
	renderTemplate(response, request, "reservation-summary.page.tmpl", &models.TemplateData { Data: data })
}

func (receiver *Repository) General(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, request, "generals.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Major(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, request, "majors.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Contact(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, request, "contacts.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) SearchAvailability(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, request, "search-availability.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) PostSearchAvailability(response http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	response.Write([]byte(fmt.Sprintf("Start date is %s, end date is %s", start, end)))
	// stringMap := make(map[string]string)
	// renderTemplate(response, request, "search-availability.page.tmpl", &models.TemplateData { StringMap: stringMap })
}

func (receiver *Repository) PostSearchAvailabilityJson(response http.ResponseWriter, request *http.Request) {
	resp := jsonResponse { OK: true, Message: "Available!" }
	out, err := json.Marshal(resp)
	if err != nil {
		helpers.ServerError(response, err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(fmt.Sprintf("%s", out)))
}

func addDefaultData(templateDataPointer *models.TemplateData, request *http.Request) *models.TemplateData {
	templateDataPointer.CSRFToken = nosurf.Token(request)
	templateDataPointer.Flash = appConfigP.Session.PopString(request.Context(), "flash")
	templateDataPointer.Error = appConfigP.Session.PopString(request.Context(), "error")
	templateDataPointer.Warning = appConfigP.Session.PopString(request.Context(), "warning")
	return templateDataPointer
}

func renderTemplate(response http.ResponseWriter, request *http.Request, templateName string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	if appConfigP.UseCache {
		templateCache = appConfigP.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	cachedTemplate, exists := templateCache[templateName]
	if !exists { log.Fatal("Could not get template cache")}
	buffer := new(bytes.Buffer)
	templateData = addDefaultData(templateData, request)
	err := cachedTemplate.Execute(buffer, templateData)
	if err != nil { helpers.ServerError(response, err) }
	_, err = buffer.WriteTo(response)
	if err != nil { helpers.ServerError(response, err) }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Println(err)
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page) // returs last element from '/'
		parsedTemplatePointer, err := template.New(name).ParseFiles(page)
		if err != nil { return cache, err }
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil { return cache, err }
		if len(layouts) > 0 {
			parsedTemplatePointer, err = parsedTemplatePointer.ParseGlob("./templates/*.layout.tmpl")
			if err != nil { return cache, err }
		}
		cache[name] = parsedTemplatePointer
	}
	return cache, nil
}
