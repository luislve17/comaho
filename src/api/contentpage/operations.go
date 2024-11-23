package contentpage

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "os"
	// "path/filepath"
	// "regexp"
	// "strings"

	"github.com/luislve17/comaho/utils"
)

func ServeContentPage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dashboardData := getContentPageData()
		err := utils.RenderTemplate(w, tmpl, "content-index", dashboardData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}

func getContentPageData() ContentPageData {
	result := ContentPageData{
		ImgURL: "foo.com",
		Name:   "FOO",
		Author: "foo fooFace",
	}
	return result
}
