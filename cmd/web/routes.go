package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/",http.HandlerFunc(app.home))
	mux.Get("/snippets/create",http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippets/create",http.HandlerFunc(app.createSnippet))
	mux.Get("/snippets/:id",http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Get("/static/",http.StripPrefix("/static",fileServer))

	return alice.New(app.recoverPanic,app.logRequest,secureHeaders).Then(mux)
}
