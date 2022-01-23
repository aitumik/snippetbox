package main

import "net/http"

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add the headers here
		w.Header().Set("X-Frame-Options","Deny")
		w.Header().Set("X-XSS-Protection","1; mode=block")
		next.ServeHTTP(w,r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log the requests here
		app.infoLogger.Printf("%s - %s %s %s",r.RemoteAddr,r.Proto,r.Method,r.URL)
		next.ServeHTTP(w,r)
	})
}
