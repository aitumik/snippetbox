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

	// Add the five routes
	mux.Get("/user/signup",dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup",dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login",dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login",dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout",dynamicMiddleware.ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Get("/static/",http.StripPrefix("/static",fileServer))

	return standardMiddleware.Then(mux)
}
