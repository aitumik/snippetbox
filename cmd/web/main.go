package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"github.com/aitumik/snippetbox/pkg"
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/aitumik/snippetbox/pkg/models/mysql"
	sqlite2 "github.com/aitumik/snippetbox/pkg/models/sqlite"
	"github.com/golangcollege/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type contextKey string

var contextKeyUser = contextKey("user")

type application struct {
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	session *sessions.Session
	cfg           *pkg.Config
	snippet       mysql.SnippetModel
	templateCache map[string]*template.Template
	users *sqlite2.UserModel
}

func main() {

	cfg := new(pkg.Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Network Address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.SecretKey,"secret","s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge","Secret Key")
	flag.Parse()

	// create loggers
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	conn, err := gorm.Open(sqlite.Open("snippet.db"), &gorm.Config{})

	db, err := conn.DB()
	if err != nil {
		errorLogger.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			errorLogger.Fatal(err)
		}
	}(db)

	// Initialize a new template cache
	templateCache, err := NewTemplateCache("./ui/html/")
	if err != nil {
		errorLogger.Fatal(err)
	}
	infoLogger.Print("Initializing the template cache")

	// Create a new session manager using the `New` method passing the secret key
	// as the parameter
	session := sessions.New([]byte("s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	// Add the templateCache to the application dependencies
	app := &application{
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
		session: session,
		cfg:         cfg,
		snippet: mysql.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
		users: &sqlite2.UserModel{
			DB: db,
		},
	}

	// Do the auto migration
	conn.AutoMigrate(&models.Snippet{})
	infoLogger.Print("Migrating database models")

	mux := app.routes()

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{tls.X25519,tls.CurveP256},
	}

	server := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLogger,
		Handler:  mux,
		TLSConfig: tlsConfig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLogger.Printf("Server started at %s", cfg.Addr)
	err = server.ListenAndServeTLS("./tls/cert.pem","./tls/key.pem")
	errorLogger.Fatal(err)
}
