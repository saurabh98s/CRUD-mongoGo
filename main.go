package main

import (
	"fmt"
	"html/template"

	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

// FormData represents the data sent to the Database
type FormData struct {
	Email    string
	DropDown string
	TextArea string
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	dropdown := r.FormValue("dropdown")
	textArea := r.FormValue("text-area")

	formData := FormData{
		Email:    email,
		DropDown: dropdown,
		TextArea: textArea,
	}
	fmt.Printf("Found Data %v", formData)
}
