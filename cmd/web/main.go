package main

import (
	"flag"
	"github.com/aitumik/snippetbox/pkg"
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/aitumik/snippetbox/pkg/models/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	cfg           *pkg.Config
	snippet       mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	cfg := new(pkg.Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Network Address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// create loggers
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	conn, err := gorm.Open(sqlite.Open("snippet.db"), &gorm.Config{})

	db, err := conn.DB()
	if err != nil {
		errorLogger.Fatal(err)
	}

	defer db.Close()

	// Initialize a new template cache
	templateCache, err := NewTemplateCache("./ui/html/")
	if err != nil {
		errorLogger.Fatal(err)
	}
	infoLogger.Print("Initializing the template cache")

	// Add the templateCache to the application dependencies
	app := &application{
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
		cfg:         cfg,
		snippet: mysql.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
	}

	// Do the auto migration
	conn.AutoMigrate(&models.Snippet{})
	infoLogger.Print("Migrating database models")

	mux := app.routes()

	infoLogger.Printf("Server started at %s", cfg.Addr)

	server := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLogger,
		Handler:  mux,
	}
	err = server.ListenAndServe()
	errorLogger.Fatal(err)
}
