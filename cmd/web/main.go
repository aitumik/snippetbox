package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aitumik/snippetbox/pkg"
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/aitumik/snippetbox/pkg/models/mysql"
	qlite "github.com/aitumik/snippetbox/pkg/models/sqlite"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golangcollege/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type contextKey string

var contextKeyUser = contextKey("user")

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	session     *sessions.Session
	cfg         *pkg.Config
	snippet     interface {
		Insert(title, content, expires string) (int, error)
		Get(id int) (*models.Snippet, error)
		Latest() ([]*models.Snippet, error)
	}
	templateCache map[string]*template.Template
	users         interface {
		Insert(name, email, password string) error
		Authenticate(email, password string) (int, error)
		Get(id int) (*models.User, error)
	}
}

func main() {

	cfg := new(pkg.Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Network Address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.SecretKey, "secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")
	flag.Parse()

	// create loggers
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	// add configurations due to docker

	// TODO move this to environment variables
	esconfig := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9300",
		},
	}
	es, err := elasticsearch.NewClient(esconfig)
	if err != nil {
		errorLogger.Fatal(err)
	}

	// Test elasticsearch connnection
	res, err := es.Info()
	if err != nil {
		errorLogger.Fatal(err)
	}

	// close the reader to avoid memory leaks
	defer res.Body.Close()

	// TODO change to postgres
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
		session:     session,
		cfg:         cfg,
		snippet: &mysql.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
		users: &qlite.UserModel{
			DB: db,
		},
	}

	// Do the auto migration
	conn.AutoMigrate(&models.Snippet{})
	infoLogger.Print("Migrating database models")

	mux := app.routes()

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	server := &http.Server{
		Addr:         cfg.Addr,
		ErrorLog:     errorLogger,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLogger.Printf("Server started at %s", cfg.Addr)
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLogger.Fatal(err)
}
