package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/ccesarfp/bookings/pkg/config"
	"github.com/ccesarfp/bookings/pkg/handlers"
	"github.com/ccesarfp/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const addr string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Starting application on port", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
