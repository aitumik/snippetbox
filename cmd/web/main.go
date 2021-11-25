package main

import (
	"flag"
	"github.com/aitumik/snippetbox/pkg"
	"log"
	"net/http"
)

func main() {

	cfg := new(pkg.Config)
	flag.StringVar(&cfg.Addr,"addr",":4000","HTTP Network Address")
	flag.StringVar(&cfg.StaticDir,"static-dir","./ui/static","Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)

	fileServer := http.FileServer(http.Dir(cfg.StaticDir))

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	log.Printf("Server started at %s",cfg.Addr)

	err := http.ListenAndServe(cfg.Addr,mux)
	log.Fatal(err)
}