package api

import (
	"html/template"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, t *template.Template) {
	mux.HandleFunc("GET /dashboard", serveDashboard(t))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
}
