package main

import (
	"fmt"
	"go_webapp/pkg/config"
	"go_webapp/pkg/constants"
	"go_webapp/pkg/handlers"
	"go_webapp/pkg/render"
	"log"
	"net/http"
)

// main serves the app
func main() {
	var app config.AppConfig

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
