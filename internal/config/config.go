package config

import(
	"html/template"
	"log"
	"github.com/alexedwards/scs/v2"
	// "io/ioutil"
	// "gopkg.in/yaml.v3"
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

// func loadConfig() {
// 	// yfile, err := ioutil.ReadFile("database.yaml")
// 	// if err != nil { log.Fatal(err) }
// 	// data := make(map[string]User)
// 	// err2 := yaml.Unmarshal(yfile, &data)
// 	// if err2 != nil { log.Fatal(err2) }
// }
