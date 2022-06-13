package handlers

import (
	"github.com/mateusjunges/lets-go/pkg/config"
	"github.com/mateusjunges/lets-go/pkg/models"
	"github.com/mateusjunges/lets-go/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.SessionManager.Put(request.Context(), "remote_ip", remoteIp)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler.
func (m *Repository) About(w http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIp := m.App.SessionManager.GetString(request.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
