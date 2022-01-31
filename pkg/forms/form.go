package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("[a-zA-Z]+@[a-zA-Z]+\\.[a-zA-Z0-9]+")

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

// Required check if the field is empty or not
func (f *Form) Required(fields ...string) {
	for _,field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field,"This field cannot be blank")
		}
	}
}

// MaxLength Implement a MaxLength method to check that a specific field in the form
// doesn't exceed the required characters
func (f *Form) MaxLength(field string,l int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > l {
		f.Errors.Add(field,fmt.Sprintf("This field is too long(maximum is %d)",l))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _,opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field,"This field is invalid")
}

func (f *Form) MinLength(field string, l int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) < l {
		f.Errors.Add(field,fmt.Sprintf("This field is too short(minimum is %d)",l))
	}
}

func (f *Form) MatchesPattern(field string,pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field,"This field is invalid")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}