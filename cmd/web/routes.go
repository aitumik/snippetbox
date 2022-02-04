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
	dynamicMiddleware := alice.New(app.session.Enable,noSurf,app.authenticate)

	mux := pat.New()
	mux.Get("/",dynamicMiddleware.ThenFunc(app.home))
	// protect the create snippets routes
	mux.Get("/snippets/create",dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippets/create",dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))

	mux.Get("/snippets/:id",dynamicMiddleware.ThenFunc(app.showSnippet))

	// Add the authentication routes
	mux.Get("/user/signup",dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup",dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login",dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login",dynamicMiddleware.ThenFunc(app.loginUser))

	// protect logout route. No point of logging out a user that is not logged in
	mux.Post("/user/logout",dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// TODO add ping handler
	mux.Get("/ping",http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir(app.cfg.StaticDir))

	// handle wildcard route
	mux.Get("/static/",http.StripPrefix("/static",fileServer))

	return standardMiddleware.Then(mux)
}
