package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mwdowns/go_webapp/pkg/config"
	"github.com/mwdowns/go_webapp/pkg/constants"
	"github.com/mwdowns/go_webapp/pkg/handlers"
	"github.com/mwdowns/go_webapp/pkg/render"

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
