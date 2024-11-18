package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/luislve17/comaho/utils"
)

func serveDashboard(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dashboardData := getDashboardData()
		err := utils.RenderTemplate(w, tmpl, dashboardData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
