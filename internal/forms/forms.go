package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if the form field is included in the submitted data and is not empty
func (f *Form) Required(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	return !(x == "")
}
