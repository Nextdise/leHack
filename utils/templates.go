package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template


func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{})  {
	err6 := templates.ExecuteTemplate(w, tmpl, data)
	fmt.Println(err6)

}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		http.ServeFile(w, r, "./templates/404.html")
	} else if status == http.StatusBadRequest {
		http.ServeFile(w, r, "./templates/400.html")
	} else {
		http.ServeFile(w, r, "./templates/500.html")
	}
}