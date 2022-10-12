package config

import(
	"html/template"
	"log"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	ErrorLog *log.Logger
	PortNumber string
	Env string
	Session *scs.SessionManager
}
