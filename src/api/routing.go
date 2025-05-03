package api

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/luislve17/comaho/api/contentpage"
	"github.com/luislve17/comaho/api/dashboard"
	"github.com/luislve17/comaho/api/processing"
)

func RegisterRoutes(r *mux.Router, t *template.Template) {
	r.HandleFunc("/dashboard", dashboard.ServeDashboard(t)).Methods("GET")
	r.HandleFunc("/{name}", contentpage.ServeContentPage(t)).Methods("GET")
	r.HandleFunc("/{name}/convert/{item}", processing.ConvertContent()).Methods("GET")
	r.HandleFunc("/{name}/convert/{item}/check", processing.CheckConvertedContent()).Methods("GET")
	r.HandleFunc("/{name}/download/{item}", processing.DownloadConvertedContent()).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("media/"))))
}
