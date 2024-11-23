package utils

import (
	"html/template"
	"net/http"
)

func ParseTemplates(tmplPath string) (*template.Template, error) {
	return template.New("templates").ParseGlob(tmplPath)
}

func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, source string, data interface{}) error {
	err := tmpl.ExecuteTemplate(w, source, data)
	return err
}
