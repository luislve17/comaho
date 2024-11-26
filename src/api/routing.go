package api

import (
	"html/template"
	"net/http"

	"github.com/luislve17/comaho/api/contentpage"
	"github.com/luislve17/comaho/api/dashboard"
	"github.com/luislve17/comaho/api/processing"
)

func RegisterRoutes(mux *http.ServeMux, t *template.Template) {
	mux.HandleFunc("GET /dashboard", dashboard.ServeDashboard(t))
	mux.HandleFunc("GET /{name}", contentpage.ServeContentPage(t))
	mux.HandleFunc("POST /{name}/download/{item}", processing.DownloadContent())
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	mux.Handle("GET /media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media/"))))
}
