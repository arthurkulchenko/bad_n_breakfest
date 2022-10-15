package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/asaskevich/govalidator"
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
	isAttrBlank := request.Form.Get(field) != ""
	if !isAttrBlank {
		form.Errors.Add(field, "Cannot be blank")
	}
	return isAttrBlank
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

func (form *Form) Invalid() bool {
	return len(form.Errors) > 0
}

func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := strings.TrimSpace(form.Get(field))
		if value == "" {
			form.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (form *Form) MinLen(length int, fields ...string) {
	for _, field := range fields {
		value := strings.TrimSpace(form.Get(field))
		if len(value) < length {
			form.Errors.Add(field, fmt.Sprintf("This field length is too short %d, expected %d", len(value), length))
		}
	}
}

func (form *Form) IsEmail(email string) {
	if !govalidator.IsEmail(form.Get(email)) {
		form.Errors.Add(email, "Invalid email address")
	}
}
