package main

import (
	"context"
	"fmt"
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/justinas/nosurf"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add the headers here
		w.Header().Set("X-Frame-Options","deny")
		w.Header().Set("X-XSS-Protection","1; mode=block")
		next.ServeHTTP(w,r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: true,
	})
	return csrfHandler
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.authenticatedUser(r) == nil {
			http.Redirect(w,r,"/user/login",302)
			return
		}
		// call the next handler in chain
		next.ServeHTTP(w,r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := app.session.Exists(r,"userID")
		if !exists {
			next.ServeHTTP(w,r)
			return
		}

		user,err := app.users.Get(app.session.GetInt(r,"userID"))
		if err == models.ErrNoRecord {
			app.session.Remove(r,"userID")
			next.ServeHTTP(w,r)
			return
		} else if err != nil {
			app.serverError(w,err)
			return
		}
		// add the user to the request context
		ctx := context.WithValue(r.Context(),contextKeyUser,user)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}


func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log the requests here
		app.infoLogger.Printf("%s - %s %s %s",r.RemoteAddr,r.Proto,r.Method,r.URL)
		next.ServeHTTP(w,r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection","close")
				app.serverError(w,fmt.Errorf("%s",err))
			}
		}()
		next.ServeHTTP(w,r)
	})
}
