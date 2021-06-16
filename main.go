package main

import (
	"forum/controller"
	"forum/utils"
	"net/http"
)

func main() {
	// Load static directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Load html templates
	utils.LoadTemplates("templates/*.html")

	// Load routes
	http.HandleFunc("/", controller.IndexPage)
	http.HandleFunc("/register", controller.RegisterPage)
	http.HandleFunc("/login", controller.LoginPage)
	http.HandleFunc("/Nouveau", controller.NouveauPage)
	http.HandleFunc("/Debat", controller.DebatPage)
	http.HandleFunc("/Sujet", controller.SujetPage)

	// Execute server
	http.ListenAndServe(":8080", nil)
}
