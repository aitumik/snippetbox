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


	f,err := os.OpenFile("/tmp/info.log",os.O_CREATE|os.O_RDWR,0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	// create loggers
	errorLogger := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime)
	infoLogger := log.New(f,"INFO\t",log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

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
	err = server.ListenAndServe()
	errorLogger.Fatal(err)
}