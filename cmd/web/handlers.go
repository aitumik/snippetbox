package main

import (
	"fmt"
	"github.com/aitumik/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippet.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	// create a value of struct TemplateData to hold  slice of snippets
	data := &TemplateData{Snippets: s}

	//files := []string{
	//	"./ui/html/home.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}
	//ts, err := template.ParseFiles(files...)
	//
	//if err != nil {
	//	app.errorLogger.Println(err.Error())
	//	app.serverError(w, err)
	//	return
	//}
	//
	//path := filepath.Join("./ui/html/", "*.page.tmpl")
	//app.infoLogger.Printf("This is the path I got from the `filepath.Join` %s", path)
	//
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.errorLogger.Println(err.Error())
	//	app.serverError(w, err)
	//	return
	//}

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

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	//files := []string{
	//	"./ui/html/show.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}
	app.render(w,r,"show.page.tmpl",data)
}
