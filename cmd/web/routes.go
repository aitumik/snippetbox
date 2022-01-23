package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/",app.home)
	mux.HandleFunc("/snippets",app.showSnippet)
	mux.HandleFunc("/snippets/create",app.createSnippet)

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	return app.logRequest(secureHeaders(mux))
}
