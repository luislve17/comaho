package api

import (
	"html/template"
	"net/http"

	"github.com/luislve17/comaho/api/contentpage"
	"github.com/luislve17/comaho/api/dashboard"
)

func RegisterRoutes(mux *http.ServeMux, t *template.Template) {
	mux.HandleFunc("GET /dashboard", dashboard.ServeDashboard(t))
	mux.HandleFunc("GET /{name}", contentpage.ServeContentPage(t))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
}
