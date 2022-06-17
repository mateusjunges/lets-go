package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mateusjunges/lets-go/internal/config"
	"github.com/mateusjunges/lets-go/internal/handlers"
	"github.com/mateusjunges/lets-go/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the application entrypoint
func main() {
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	app.SessionManager = session

	templateCache, err := render.RegisterTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache.")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repository := handlers.NewRepository(&app)
	handlers.NewHandlers(repository)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting web server at port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = server.ListenAndServe()

	log.Fatal("could not start server", err)
}
