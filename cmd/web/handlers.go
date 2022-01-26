package main

import (
	"fmt"
	"github.com/aitumik/snippetbox/pkg/models"
	"net/http"
	"strconv"
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
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, 405)
		return
	}

	title := "Are we alone"
	content := "Are we alone in the universe that is the question we are asking ourselves this evening"

	id, err := app.snippet.Insert(title, content, "7")
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippets?id=%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("Creating the snippet form..."))
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
