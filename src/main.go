package main

import (
	"log"
	"net/http"

	"github.com/luislve17/comaho/api"
	"github.com/luislve17/comaho/utils"
)

func main() {
	DEFAULT_PORT := "8080"
	mux := http.NewServeMux()

	// Load the template once
	tmplPath := "templates/*.html"
	parsedTmpl, err := utils.ParseTemplates(tmplPath)

	if err != nil {
		log.Printf("Server error %s...\n", err)
		return
	}

	api.RegisterRoutes(mux, parsedTmpl)

	// Start the server
	log.Printf("Server starting on port %s...\n", DEFAULT_PORT)
	server := &http.Server{
		Addr:    ":" + DEFAULT_PORT,
		Handler: mux,
	}
	log.Println("Listening...")
	log.Println(server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Server error %s...\n", err)
		return
	}
}
