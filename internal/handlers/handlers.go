package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mateusjunges/lets-go/internal/config"
	"github.com/mateusjunges/lets-go/internal/models"
	"github.com/mateusjunges/lets-go/internal/render"
	"log"
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

	render.Template(w, request, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler.
func (m *Repository) About(w http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIp := m.App.SessionManager.GetString(request.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.Template(w, request, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, request *http.Request) {
	render.Template(w, request, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Majors renders the major room page
func (m *Repository) Majors(w http.ResponseWriter, request *http.Request) {
	render.Template(w, request, "majors.page.tmpl", &models.TemplateData{})
}

// Generals renders the general quarters room page
func (m *Repository) Generals(w http.ResponseWriter, request *http.Request) {
	render.Template(w, request, "generals.page.tmpl", &models.TemplateData{})
}

// DisplaySearchAvailability renders the search availability page
func (m *Repository) DisplaySearchAvailability(w http.ResponseWriter, request *http.Request) {
	render.Template(w, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, request *http.Request) {
	render.Template(w, request, "contact.page.tmpl", &models.TemplateData{})
}

// SearchAvailability handles request for availability and renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	_, _ = w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// SearchAvailabilityJson handles request for availability and responds with JSON
func (m *Repository) SearchAvailabilityJson(w http.ResponseWriter, request *http.Request) {
	response := jsonResponse{
		Ok:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(response, "", "    ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(out)
}
