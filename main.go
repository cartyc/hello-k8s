package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/index.html"))

	t.Execute(w, nil)
}

// Health Check Endpoint for K8s
func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "healthy")
}

func metrics(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "healthy")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/healthz", healthCheck)
	r.HandleFunc("/metrics", healthCheck)

	port := ":8080"

	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, r)
}
