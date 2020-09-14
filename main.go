package main

import (
	"fmt"
	"html/template"

	"net/http"

	"github.com/gorilla/mux"
	// "github.com/globalsign/mgo"
)

var tpl *template.Template

// FormData represents the data sent to the Database
type FormData struct {
	Email    string
	Password string
	
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/index", serveTemplate)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))
	r.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("username")
	password := r.FormValue("password")

	formData := FormData{
		Email:    email,
		Password: password,
	}
	fmt.Printf("Found Data %v", formData)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "index.html")
}
