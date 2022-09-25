package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"path/filepath"
	"bytes"
	"github.com/arthurkulchenko/bed_n_breakfest/pkg/config"
	// "github.com/arthurkulchenko/bed_n_breakfest/pkg/models"
)

var appConfigP *config.AppConfig
var RepositoryPointer *Repository
// RepositoryPointer.AppConfigPointer => *config.AppConfig

type Repository struct {
	AppConfigPointer *config.AppConfig
}

func SetConfigAndRepository(appConfigPointer *config.AppConfig) {
	appConfigP = appConfigPointer
	RepositoryPointer = &Repository { AppConfigPointer: appConfigPointer }
}

func NewRepo(pointer *config.AppConfig) *Repository {
	return &Repository { AppConfigPointer: pointer }
}

func NewHandlers(repositoryPointer *Repository) {
	RepositoryPointer = repositoryPointer
}

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

func (receiver *Repository) Home(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	session := receiver.AppConfigPointer.Session
	stringMap["remoteaddr"] = request.RemoteAddr
	session.Put(request.Context(), "remoteaddr", request.RemoteAddr)

	renderTemplate(response, "home.page.tmpl", &TemplateData { StringMap: stringMap })
}

func (receiver *Repository) About(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	session := receiver.AppConfigPointer.Session
	stringMap["remoteaddr"] = session.GetString(request.Context(), "remoteaddr")

	renderTemplate(response, "about.page.tmpl", &TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Reservation(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	session := receiver.AppConfigPointer.Session
	stringMap["remoteaddr"] = session.GetString(request.Context(), "remoteaddr")

	renderTemplate(response, "reservation.page.tmpl", &TemplateData { StringMap: stringMap })
}

func (receiver *Repository) General(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, "generals.page.tmpl", &TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Major(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, "majors.page.tmpl", &TemplateData { StringMap: stringMap })
}

func (receiver *Repository) Contact(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	renderTemplate(response, "contacts.page.tmpl", &TemplateData { StringMap: stringMap })
}

func addDefaultData(templateDataPointer *TemplateData) *TemplateData {
	return templateDataPointer
}

func renderTemplate(response http.ResponseWriter, templateName string, templateData *TemplateData) {
	var templateCache map[string]*template.Template
	if appConfigP.UseCache {
		templateCache = appConfigP.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	cachedTemplate, exists := templateCache[templateName]
	if !exists { log.Fatal("Could not get template cache")}
	buffer := new(bytes.Buffer)
	templateData = addDefaultData(templateData)
	err := cachedTemplate.Execute(buffer, templateData)
	if err != nil { log.Println(err) }
	_, err = buffer.WriteTo(response)
	if err != nil { log.Println(err) }
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
