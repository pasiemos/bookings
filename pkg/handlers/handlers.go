package handlers

import (
	"github.com/pasiemos/bookings/pkg/config"
	"github.com/pasiemos/bookings/pkg/models"
	"github.com/pasiemos/bookings/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,

	}
}
//NewHandlers sets the repository for the handlers
//r is type pointer to Repository
//& a reference to Repository
func NewHandlers(r *Repository){
	Repo = r
}


//Home is the home page handler
//m is a new variable  of type pointer to Repository. It's a receiver of type Repo?
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perfrorm some logic

	stringMap :=make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP


	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl.html", &models.TemplateData{})
}

//Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl.html", &models.TemplateData{})
}

//Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl.html", &models.TemplateData{})
}

//Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl.html", &models.TemplateData{})
}

//Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl.html", &models.TemplateData{})
}