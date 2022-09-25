package models

import "github.com/arthurkulchenko/bed_n_breakfest/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Watrning string
	Error string
	Form *forms.Form
}
