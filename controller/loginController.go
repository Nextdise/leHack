package controller

import (
	"database/sql"
	"fmt"
	"forum/models"
	"forum/utils"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.URL.String() != "/login" {
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, "error.html", 400)
			return
		}
		r.ParseForm()
		// take values to send to DB
		searchEmail := r.FormValue("email")
		searchPassword := r.FormValue("pwd")

		//searchPassword = string(hash)

		database, err := sql.Open("sqlite3", "Forum.db")
		if err != nil {
			log.Fatalln()
		}
		var PassData string
		err = database.QueryRow("SELECT Password FROM Users WHERE Email=?", searchEmail).Scan(&PassData)
		if err != nil {
			panic(err)
		}
		fmt.Println(searchPassword, PassData)
		err3 := bcrypt.CompareHashAndPassword([]byte(PassData), []byte(searchPassword))
		if err3 != nil {
			fmt.Println("azezaed")

		}

		rows, err7 := database.Query("SELECT * FROM Users WHERE Email = '" + searchEmail + "'")
		if err7 != nil {
			log.Fatalln()

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

		if len(users) != 1 {
			p := models.LoginTemplate{Message: "No accout found."}
			utils.ExecuteTemplate(w, "login.html", p)
			return
		}
		fmt.Println(users)
		if err != nil {
			p := models.LoginTemplate{Message: "Wrong password."}
			utils.ExecuteTemplate(w, "login.html", p)
			return
		}

		p := models.LoginTemplate{Message: "You are successfully accessed."}
		time.Sleep(2 * time.Second)
		utils.ExecuteTemplate(w, "index2.html", p)

	} else if r.Method == "GET" {
		if r.URL.Path != "/login" {
			utils.ErrorHandler(w, r, http.StatusNotFound) //404 error
			return
		}
		p := models.LoginTemplate{}
		utils.ExecuteTemplate(w, "login.html", p)
	}
}
