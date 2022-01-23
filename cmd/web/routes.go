package main

import (
	"net/http"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/",app.home)
	mux.HandleFunc("/snippets",app.showSnippet)
	mux.HandleFunc("/snippets/create",app.createSnippet)

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	return alice.New(app.recoverPanic,app.logRequest,secureHeaders).Then(mux)
}
