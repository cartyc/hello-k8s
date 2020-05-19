package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Rancher Master Class")

}

func main() {

	http.HandleFunc("/", home)

	port := ":8080"

	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
