package handlers

import (
	"fmt"
	"go_webapp/pkg/config"
	"go_webapp/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the home page")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the about page")
	render.RenderTemplate(w, "about.page.tmpl")
}
