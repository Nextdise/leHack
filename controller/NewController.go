package controller

import (
	"database/sql"
	"fmt"
	"forum/utils"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Nom      string
	Contenue string
}
type Database struct {
	Posts []Post
}

func NouveauPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.URL.String() != "/Nouveau" {

			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}

		database, _ := sql.Open("sqlite3", "Forum.db")
		_, err4 := database.Exec("DELETE * FROM posts")
		if err4 != nil {
			fmt.Print("erreur")
			http.Error(w, "excu", 500)
			return

		}

	} else if r.Method == "GET" {

		if r.URL.Path != "/Nouveau" {

			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		//puts the database information in the card
		p, err := database.Query("SELECT Nom,Contenue FROM Posts ")
		if err != nil {
			panic(err)
		}
		var data Database
		for p.Next() {
			var pt Post
			p.Scan(&pt.Nom, &pt.Contenue)
			data.Posts = append(data.Posts, pt)
		}
		fmt.Println(data)

		utils.ExecuteTemplate(w, "Nouveau.html", data)

	}

}
