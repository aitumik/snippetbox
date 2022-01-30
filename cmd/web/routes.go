package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	// standard middleware
	standardMiddleware := alice.New(app.recoverPanic,app.logRequest,secureHeaders)
	// dynamic middleware
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/",dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippets/create",dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippets/create",dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippets/:id",dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Get("/static/",http.StripPrefix("/static",fileServer))

	return standardMiddleware.Then(mux)
}
