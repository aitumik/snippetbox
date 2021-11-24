package main

import (
	"log"
	"net/http"
)

// note that this will handle all the request
// even if you visit the http://localhost:4000/foo/bar
// it will still be handled by this handler
func home(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("Hello world from Snippetbox"));
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	log.Println("Starting server at :4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}