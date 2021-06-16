package controller

import (
	_ "github.com/mattn/go-sqlite3"

	"forum/models"
	"forum/utils"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.URL.String() != "/" {
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}

		p := models.IndexTemplate{}
		utils.ExecuteTemplate(w, "index.html", p)

	} else if r.Method == "GET" {
		if r.URL.Path != "/" {
			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		p := models.IndexTemplate{}
		utils.ExecuteTemplate(w, "index.html", p)

	}
}
