package controller

import (
	"forum/models"
	"forum/utils"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func DebatPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.URL.String() != "/Debat" {

			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}
		p := models.DebatTemplate{}
		utils.ExecuteTemplate(w, "Debat.html", p)
	} else if r.Method == "GET" {

		if r.URL.Path != "/Debat" {

			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		p := models.DebatTemplate{}
		utils.ExecuteTemplate(w, "Debat.html", p)
	}
}
