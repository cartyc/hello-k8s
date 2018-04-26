package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/index.html"))

	t.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", home)

	http.ListenAndServe(":80", nil)

}
