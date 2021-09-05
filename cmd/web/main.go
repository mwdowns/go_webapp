package main

import (
	"fmt"
	"go_webapp/pkg/config"
	"go_webapp/pkg/constants"
	"go_webapp/pkg/handlers"
	"go_webapp/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

// main serves the app
func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    constants.PortNumber,
		Handler: routes(&app),
	}

	fmt.Println("listing on port", constants.PortNumber)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("could not serve app")
	}
}
