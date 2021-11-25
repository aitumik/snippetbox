package main

import (
	"flag"
	"github.com/aitumik/snippetbox/pkg"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	cfg 		*pkg.Config
}

func main() {

	cfg := new(pkg.Config)
	flag.StringVar(&cfg.Addr,"addr",":4000","HTTP Network Address")
	flag.StringVar(&cfg.StaticDir,"static-dir","./ui/static","Path to static assets")
	flag.Parse()

	// create loggers
	infoLogger := log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr,"INFO\t",log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	app := &application{
		errorLogger: errorLogger,
		infoLogger: infoLogger,
		cfg: cfg,
	}

	mux := app.routes()

	infoLogger.Printf("Server started at %s",cfg.Addr)

	server := &http.Server{
		Addr: cfg.Addr,
		ErrorLog: errorLogger,
		Handler: mux,
	}
	err := server.ListenAndServe()
	errorLogger.Fatal(err)
}