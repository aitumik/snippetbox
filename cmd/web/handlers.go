package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Welcome to Snippetbox"))
}

func createSnippet(w http.ResponseWriter,r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow","POST")
		http.Error(w,"Method not supported",405)
	}

	w.Write([]byte("Creating a snippet..."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Fprint(w,"Selecting snippet with ID= %d...",id)
}