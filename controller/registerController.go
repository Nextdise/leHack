package controller

import (
	"database/sql"
	"fmt"
	"forum/models"
	"forum/utils"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if r.URL.String() != "/register" {
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}

		// take values to send to DB
		searchUserName := r.FormValue("username")
		searchEmail := r.FormValue("email")
		searchPassword := r.FormValue("pwd")
		//hash password
		hash, _ := bcrypt.GenerateFromPassword([]byte(searchPassword), 10)
		searchPassword = string(hash)

		database, _ := sql.Open("sqlite3", "Forum.db")
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, UserName TEXT, Email TXT, Password TXT)")
		statement.Exec()

		rows, err4 := database.Query("SELECT * FROM Users WHERE Email = '" + searchEmail + "' OR UserName = '" + searchUserName + "'")
		if err4 != nil {
			fmt.Print("erreur")
			http.Error(w, "excu", 500)
			return

		}
		var users []models.User
		// Iterate and save each row of the Query into the users table
		for rows.Next() {
			item := models.User{}
			err2 := rows.Scan(&item.ID, &item.UserName, &item.Email, &item.Password)
			if err2 != nil {
				panic(err2)

			}
			users = append(users, item)
		}

		if len(users) != 0 {
			p := models.RegisterTemplate{Message: "Account already exist."}
			utils.ExecuteTemplate(w, "register.html", p)
			return
		}

		statement, _ = database.Prepare("INSERT INTO Users (UserName, Email, Password) VALUES (?, ?, ?)")
		statement.Exec(searchUserName, searchEmail, searchPassword)

		p := models.RegisterTemplate{Message: "Registered successfully."}
		utils.ExecuteTemplate(w, "register.html", p)

	} else if r.Method == "GET" {
		if r.URL.Path != "/register" {
			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		p := models.RegisterTemplate{}
		utils.ExecuteTemplate(w, "register.html", p)
	}
}
