package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandeler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v\n", err)
		return
	}
	fmt.Fprintf(w, "Post info successfully\n")
	email := r.FormValue("email")
	username := r.FormValue("username")
	fmt.Fprintf(w, "your email = %v\nyour username = %v\n", email, username)
}

func greet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Welcome to a response in golang")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", greet)
	http.HandleFunc("/formHandeler", formHandeler)

	fmt.Println("Starting golang Web Server at :3000 port")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
