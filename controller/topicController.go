package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"forum/models"
	"forum/utils"
	"net/http"
)

var database, _ = sql.Open("sqlite3", "Forum.db")

func SujetPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		if r.URL.String() != "/Sujet" {

			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}
		searchNomContenue := r.FormValue("NomContenue")
		searchContenue := r.FormValue("Contenue")

		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Nom TEXT, Contenue TXT)")
		statement.Exec()

		rows, err4 := database.Query("SELECT * FROM Posts WHERE Nom = '" + searchNomContenue + "' OR Contenue = '" + searchContenue + "'")
		if err4 != nil {
			fmt.Print("erreur")
			http.Error(w, "excu", 500)
			return

		}
		var Posts []models.Posts
		// Iterate and save each row of the Query into the users table
		for rows.Next() {
			item := models.Posts{}
			err2 := rows.Scan(&item.ID, &item.Nom, &item.Contenue)
			if err2 != nil {
				panic(err2)

			}
			Posts = append(Posts, item)
		}
		statement, _ = database.Prepare("INSERT INTO Posts (Nom, Contenue) VALUES (?, ?)")
		statement.Exec(searchNomContenue, searchContenue)

		p := models.SujetTemplate{Message: "Registered successfully."}
		utils.ExecuteTemplate(w, "topic.html", p)

	} else if r.Method == "GET" {

		if r.URL.Path != "/Sujet" {

			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		p := models.SujetTemplate{}
		utils.ExecuteTemplate(w, "topic.html", p)
	}
}
