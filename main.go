package main

import (
	"fmt"
	"html/template"
	"log"
	"mongo-golang/config"
	"mongo-golang/persistent"
	"mongo-golang/persistent/models"
	"mongo-golang/service"
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

	var mongoConfig config.Mongo

	mongoConfig.Load()
	persistent.Init()

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
	fmt.Println("inside hello")
	email := r.FormValue("username")
	password := r.FormValue("password")

	var formData *models.User

	formData = &models.User{
		Email:    email,
		Password: password,
	}
	if formData != nil {
		fmt.Printf("Found Data %v", formData)
		err := service.Init().Auth().Save(formData)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Fatalln("[ERROR] No data Found")

}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "index.html")
}
