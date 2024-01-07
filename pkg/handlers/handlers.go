package handlers

import (
	"github.com/ccesarfp/bookings/pkg/config"
	"github.com/ccesarfp/bookings/pkg/models"
	"github.com/ccesarfp/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

const posfix = ".page.tmpl"

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(writer, "home"+posfix, &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about"+posfix, &models.TemplateData{
		StringMap: stringMap,
	})
}
