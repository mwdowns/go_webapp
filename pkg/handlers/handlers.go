package handlers

import (
	"fmt"
	"go_webapp/pkg/config"
	"go_webapp/pkg/models"
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

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the about page")

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	strMap := make(map[string]string)
	strMap["test"] = "Heyo"
	strMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: strMap})
}
