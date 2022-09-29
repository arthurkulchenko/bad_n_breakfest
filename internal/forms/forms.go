package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (form *Form) Has(field string, request *http.Request) bool {
	isAttrPresent := request.Form.Get(field) != ""
	if !isAttrPresent {
		form.Errors.Add(field, "Cannot be blank")
	}
	return isAttrPresent
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}
