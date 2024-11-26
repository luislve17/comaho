package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luislve17/comaho/api"
	"github.com/luislve17/comaho/utils"
)

func main() {
	DEFAULT_PORT := "8080"
	r := mux.NewRouter()

	tmplPath := "templates/*.html"
	parsedTmpl, err := utils.ParseTemplates(tmplPath)

	if err != nil {
		log.Printf("Server error %s...\n", err)
		return
	}

	api.RegisterRoutes(r, parsedTmpl)

	// Start the server
	log.Printf("Server starting on port %s...\n", DEFAULT_PORT)
	server := &http.Server{
		Addr:    ":" + DEFAULT_PORT,
		Handler: r,
	}
	log.Println("Listening...")
	log.Println(server.Addr)
	log.Fatal(server.ListenAndServe())
}
