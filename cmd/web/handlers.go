package main

import (
	"fmt"
	"github.com/aitumik/snippetbox/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// pat now handles request to the `/` route
	s, err := app.snippet.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	// create a value of struct TemplateData to hold  slice of snippets
	data := &TemplateData{Snippets: s}
	app.render(w,r,"home.page.tmpl",data)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// First wee need to call the method r.ParseForm() which loads the values of the post
	// to the r.PostForm map
	// We can get for example title if we do this `r.ParseForm().Get("Title")`
	// Note that the r.ParseForm() is limited to 10MB
	// To change this limit use the http.MaxBytesReader()
	r.Body = http.MaxBytesReader(w,r.Body,4096)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w,http.StatusBadRequest)
		return
	}

	// load them in variables : you could use them directly also
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	// Validate the title field
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "The field is too long(maximum is 100 characters)"
	}

	// Validate the content field
	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	// Validate the expires field
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	// If there are any errors dump them
	if len(errors) > 0 {
		// pass back the errors and the url.Values type which is a map
		data := &TemplateData{
			FormErrors: errors,
			FormData: r.PostForm,
		}
		app.render(w,r,"create.page.tmpl",data)
		return
	}

	id,err := app.snippet.Insert(title,content,expires)
	if err != nil {
		app.serverError(w,err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter,r *http.Request) {
	app.render(w,r,"create.page.tmpl",nil)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippet.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	// Create and instance of a TemplateData struct holding the snippet data
	data := &TemplateData{Snippet: s}

	app.render(w,r,"show.page.tmpl",data)
}
