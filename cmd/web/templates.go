package main

import (
	"github.com/aitumik/snippetbox/pkg/models"
	"html/template"
	"net/url"
	"path/filepath"
	"time"
)

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
	CurrentYear int
	FormData url.Values
	FormErrors map[string]string
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages,err := filepath.Glob(filepath.Join(dir,"*.page.tmpl"))
	if err != nil {
		return nil,err
	}
	for _, page := range pages {
		// Extract the file name(like 'home.page.tmpl') from the full path
		// and assign it to the name variable
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before the ParseFiles
		// method
		ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil,err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil,err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil,err
		}

		cache[name] = ts
	}

	return cache, nil
}

// humanDate function returns a nicely formatted string representation of
// time.Time value
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize the template.FuncMap value with the string-keyed map
// which acts as a lookup table
var functions = template.FuncMap{
	"humanDate": humanDate,
}
